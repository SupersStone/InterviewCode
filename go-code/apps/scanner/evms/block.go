package evms

import (
	"dao-exchange/apps/scanner/repo"
	"dao-exchange/internal/models/scanner"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

const maxReceipt = 10

// BlockScanner scanner evm like block
type BlockScanner struct {
	chain     string
	chainID   int
	client    *ethclient.Client
	rpcClient *rpc.Client
}

// 处理区块里的交易receipt并更新数据库高度
type rpcBlock struct {
	Hash         common.Hash    `json:"hash"`
	Transactions []common.Hash  `json:"transactions"`
	Number       hexutil.Uint64 `json:"number"`
	Timestamp    hexutil.Uint64 `json:"timestamp"`
}

// New BlockScanner
func New(chain string, chainID int, client *ethclient.Client, rpcClient *rpc.Client) *BlockScanner {
	return &BlockScanner{
		chain:     chain,
		chainID:   chainID,
		client:    client,
		rpcClient: rpcClient,
	}
}

// TaskName build taskName
func (s *BlockScanner) TaskName() string {
	return s.chain + ":" + strconv.Itoa(s.chainID)
}

// BatchHandleBlock batch handle block
func (s *BlockScanner) BatchHandleBlock(blockNumbers []uint64) (recipets []*ScanLogsPtr, err error) {
	var (
		elems  []rpc.BatchElem
		blocks = make([]*rpcBlock, len(blockNumbers))
	)

	blocks, err = getBlockByNumbers(s.rpcClient, blockNumbers)
	if err != nil {
		return nil, err
	}

	for i, block := range blocks {
		if block == nil {
			return nil, errors.Errorf("get block: %d failed, err: %v", blockNumbers[i], elems[i].Error.Error())
		}

		res, err := handleTxsInBlock(s.rpcClient, block)
		if err != nil {
			return nil, errors.Wrapf(err, "handle block number: %d failed, block: %+v", block.Number, block)
		}

		recipets = append(recipets, res)
	}

	// 成功处理完一批blocks之后更新数据库扫块记录
	return recipets, nil
}

// SyncToDB write to db
func (s *BlockScanner) SyncToDB(dao *repo.Dao, blockHeight uint64) (err error) {
	if err = dao.UpdateBlockHeight(s.TaskName(), s.chain, &scanner.BlockScanHeight{
		Height: blockHeight,
	}); err != nil {
		err = errors.Wrapf(err, "update handled block number failed, handled block: %d, task name: %s, chain: %s", blockHeight, s.TaskName(), s.chain)
	}

	return
}
