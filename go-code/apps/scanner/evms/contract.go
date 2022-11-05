/*
When we add a new contract, it's start height is lower than current block height.
we should begin a new go runtine to handle scan log.
*/

package evms

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"dao-exchange/internal/models/scanner"
	"dao-exchange/internal/mqs"
	"dao-exchange/internal/utils"
	"dao-exchange/pkg/retry"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	initScanPrefix  = "init:contract:scan"
	initScanTimeout = time.Minute
	timeout         = time.Second * 5
)

// ContractScanner scanner contract
type ContractScanner struct {
	chain      string
	chainID    int
	topic      string
	delayBlock uint64
	client     *ethclient.Client
	rpcClient  *rpc.Client
	rdb        *redis.ClusterClient
	producer   *mqs.Producer
}

// FilterLogs filter log
func (s *ContractScanner) FilterLogs(startBlock, endBlock, maxScanElem uint64, contract *scanner.BlockScanContract, stop chan struct{}) {
	for i := startBlock; i < endBlock; i = i + maxScanElem {
		select {
		case <-stop:
			logrus.Infof("init contract %s scan got exit signal, exit", contract.Address)
			return
		default:
			end := startBlock + maxScanElem
			for {
				if err := s.filterLog(startBlock, end, contract); err == nil {
					continue
				}
			}
		}
	}
}

// FilterLog filter log
func (s *ContractScanner) filterLog(startBlock, endBlock uint64, contract *scanner.BlockScanContract) error {
	var (
		logs []types.Log
		err  error
	)
	if err = retry.BackoffRetry(
		func() error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			logs, err = s.client.FilterLogs(
				ctx,
				ethereum.FilterQuery{
					FromBlock: big.NewInt(0).SetUint64(startBlock),
					ToBlock:   big.NewInt(0).SetUint64(endBlock),
					Addresses: []common.Address{common.HexToAddress(contract.Address)},
				},
			)
			logrus.Warnf("init scan contract: %s get logs from: %d to: %d block err: %s", contract.Address, startBlock, endBlock, err)
			return err
		},
	); err != nil {
		return err
	}

	blockNumbers := utils.NewSlice(startBlock, endBlock-startBlock, endBlock, 1)
	blocks, err := getBlockByNumbers(s.rpcClient, blockNumbers)
	if err != nil {
		return err
	}

	blockMap := make(map[string]uint64)
	for _, block := range blocks {
		blockMap[block.Hash.Hex()] = uint64(block.Timestamp)
	}

	for _, log := range logs {
		msg := &Msg{
			Log:            &log,
			ChainID:        s.chainID,
			ChainName:      s.chain,
			BlockTimestamp: blockMap[log.BlockHash.Hex()],
		}
		data, err := json.Marshal(msg)
		if err != nil {
			err = errors.Wrapf(err, "marshal %v failed", msg)
			logrus.Warn(err)
			continue
		}

		s.producer.PublishMsg(log.Address.Bytes(), data)
	}
	logrus.Infof("init scan contract: %s logs len in %d - %d : %d", contract.Address, startBlock, endBlock, len(logs))

	return nil
}
