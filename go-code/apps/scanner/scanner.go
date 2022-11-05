package scanner

import (
	"context"
	"strconv"
	"sync"
	"time"

	"dao-exchange/apps/scanner/evms"
	"dao-exchange/apps/scanner/repo"
	"dao-exchange/config"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models/scanner"
	"dao-exchange/internal/mqs"
	"dao-exchange/internal/utils"
	"dao-exchange/pkg/myerr"
	"dao-exchange/pkg/retry"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// default value
const (
	TaskName       = "chainScanner"
	timeout        = time.Second * 5
	maxElem        = 20
	maxReceipt     = 10
	defaultTryTime = 10
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// Scanner name
type Scanner struct {
	sync.Mutex
	blockScanner    *evms.BlockScanner
	chain           string
	chainID         int
	topic           string
	client          *ethclient.Client
	rpcClient       *rpc.Client
	rdb             cache.Client
	producer        *mqs.Producer
	repo            *repo.Dao
	cache           *ContractsCache
	delayBlock      uint64
	startBlock      uint64
	scanAmount      int
	scanInterval    int
	initScanAddress string
	initScan        chan *scanner.BlockScanContract
	stop            chan struct{}
}

// New new sacnner
func New(chain string, chainID int, rpc *rpc.Client, rdb cache.Client, dao *repo.Dao, producer *mqs.Producer, options ...Option) *Scanner {
	cli := ethclient.NewClient(rpc)
	blockScanner := evms.New(chain, chainID, cli, rpc)
	s := Scanner{
		chain:        chain,
		chainID:      chainID,
		stop:         make(chan struct{}),
		client:       cli,
		rpcClient:    rpc,
		rdb:          rdb,
		repo:         dao,
		producer:     producer,
		initScan:     make(chan *scanner.BlockScanContract, 1),
		cache:        NewContractsCache(chain, chainID, dao),
		blockScanner: blockScanner,
	}

	for _, o := range options {
		o(&s)
	}

	return &s
}

// Start scanner
func Start(conf *config.Config, rdb cache.Client, dao *repo.Dao, stopCh <-chan struct{}) (err error) {
	for _, scannerConf := range conf.Scanner {
		// init rpc client
		logrus.Infof("start chain scanner: %s, node: %s", scannerConf.Chain, scannerConf.NodeURL)
		cli, err := rpc.DialContext(context.Background(), scannerConf.NodeURL)
		if err != nil {
			return errors.Wrapf(err, "connect to node: %s failed", scannerConf.NodeURL)
		}

		options := []Option{
			WithDelayBlock(scannerConf.DelayBlock),
			WithScanInterval(scannerConf.ScanInterval),
			WithScanAmount(scannerConf.ScanAmount),
			WithStartBlock(scannerConf.StartBlock),
			WithChainID(scannerConf.ChainID),
		}
		scanner := New(
			scannerConf.Chain,
			scannerConf.ChainID,
			cli,
			rdb,
			dao,
			mqs.NewProducerWithCerd(conf.ProducerKafka, conf.Credential.AccessKeyID, conf.Credential.SecretAccessKey),
			options...,
		)

		go func(s *Scanner) {
			s.run(stopCh)
		}(scanner)
	}
	return nil
}

func (s *Scanner) buildLockNmae() string {
	return TaskName + ":" + s.chain + ":" + strconv.Itoa(s.chainID)
}

func (s *Scanner) run(stopCh <-chan struct{}) {
	// cancel方法用来释放续锁的守护goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer myerr.Recover(cancel)
	defer cancel()

	lockID := uuid.New()
	lockName := s.buildLockNmae()
	// 尝试获取锁，lockID用来区分是不是自己持有的锁
	cache.TryLock(ctx, lockName, lockID.String(), 30*time.Second, s.rdb)

	// 获取保存的合约地址
	if err := s.cache.CacheContracts(); err != nil {
		logrus.Error(err)
	}

	// 定时更新表里的合约地址
	go s.cache.StartCacheContracts(stopCh, time.Second*30)

	for {
		select {
		case <-stopCh:
			logrus.Infof("%s get cancel signal, exit", lockName)
			close(s.stop)
			s.Close()
			return
		default:
			if err := s.scanBlock(); err != nil {
				logrus.Error(err)
			}
			time.Sleep(time.Second * time.Duration(s.scanInterval))
		}
	}
}

func (s *Scanner) buildTaskName(chain string, chainID int) string {
	return s.chain + ":" + strconv.Itoa(s.chainID)
}

func (s *Scanner) getCurrentDBHeight(tryTime int) (currentDbHeight uint64, err error) {
	taskName := s.buildTaskName(s.chain, s.chainID)
	err = retry.BackoffRetry(
		func() error {
			currentDbHeight, err = s.repo.QueryCurrentHeightWithTaskName(taskName)
			if err == nil {
				return nil
			}

			if errors.Is(err, myerr.ErrRecordNotFound) {
				logrus.Infof("%s no scan record in db, insert start block: %d to db", taskName, s.startBlock)
				if err = s.repo.Insert(&scanner.BlockScanHeight{
					Height:   s.startBlock,
					Chain:    s.chain,
					TaskName: taskName,
				}); err != nil {
					return errors.Wrapf(err, "chain: %s insert start block height to db failed", taskName)
				}
				currentDbHeight = s.startBlock
				return nil
			}
			logrus.Warnf("query chain: %s current scan block failed, err: %s", taskName, err.Error())
			return err
		},
	)

	return
}

func getBlockWithTry(client *ethclient.Client, tryTime int) (latestHeight uint64, err error) {
	err = retry.BackoffRetry(
		func() error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			latestHeight, err = client.BlockNumber(ctx)
			cancel()

			return err
		},
	)

	return
}

func (s *Scanner) scanBlock() (err error) {
	defer myerr.Recover(nil)
	var (
		currentDbHeight uint64
		latestHeight    uint64
	)
	taskName := s.buildTaskName(s.chain, s.chainID)
	currentDbHeight, err = s.getCurrentDBHeight(defaultTryTime)
	if err != nil {
		return errors.Wrapf(err, "%s query db current height failed", taskName)
	}

	latestHeight, err = getBlockWithTry(s.client, defaultTryTime)
	if err != nil {
		return errors.Wrapf(err, "%s get current block number failed, error: %s", taskName, err)
	}

	logrus.Infof("%s current db height is: %d, current block height is: %d", taskName, currentDbHeight, latestHeight)
	// 未达到延迟块
	if latestHeight-currentDbHeight-1 < s.delayBlock {
		logrus.Infof("%s: current db scanned block: %d, latest block: %d, not reach dealy block: %d, wait next loop", taskName, currentDbHeight, latestHeight, s.delayBlock)
		return
	}

	// 处理区块中log
	startHeight := currentDbHeight + 1
	endHeight := latestHeight - s.delayBlock
	for i := startHeight; i <= endHeight; i += maxElem {
		select {
		case <-s.stop:
			logrus.Infof("%s scan block get exit signal", taskName)
			return
		default:
			blocks := utils.NewSlice(i, maxElem, endHeight, 1)
			if err = s.batchHandleBlock(blocks); err != nil {
				return errors.Wrapf(err, "batch handle block from: %d to: %d failed", blocks[0], i)
			}
			continue
		}
	}

	return
}

func (s *Scanner) batchHandleBlock(blockNumbers []uint64) (err error) {
	blockReceipts, err := s.blockScanner.BatchHandleBlock(blockNumbers)
	if err != nil {
		return err
	}

	if err := s.blockReceipts(blockReceipts); err != nil {
		return err
	}

	return s.blockScanner.SyncToDB(s.repo, blockNumbers[len(blockNumbers)-1])
}

func (s *Scanner) blockReceipts(resp []*evms.ScanLogsPtr) (err error) {
	if len(resp) <= 0 {
		return nil
	}

	for _, receipts := range resp {
		if err := s.txReceipts(receipts); err != nil {
			return err
		}
	}

	return
}

func (s *Scanner) txReceipts(receipts *evms.ScanLogsPtr) (err error) {
	for _, log := range receipts.Logs {
		// TODO 是支持一整个交易还是支持单个log发送，判断是否是支持的合约
		if !s.cache.Contains(log.Address) {
			continue
		}

		// 是支持的合约事件 发送到Kafka
		msg := &evms.Msg{
			Log:            log,
			ChainID:        s.chainID,
			ChainName:      s.chain,
			BlockTimestamp: receipts.Timestamp,
		}
		data, err := json.Marshal(msg)
		if err != nil {
			return errors.Wrapf(err, "marshal %v failed", msg)
		}

		if err := s.producer.PublishMsg(log.Address.Bytes(), data); err != nil {
			s.producer = s.producer.NewWriterWithCerd()
			return errors.Wrapf(err, "publish msg failed")
		}
	}
	return
}

// Close when quit
func (s *Scanner) Close() {
	s.client.Close()
	s.repo.Close()
	s.rdb.Close()
	s.producer.Close()
}
