package evms

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

func getBlockByNumbers(rpcClient *rpc.Client, blockNumbers []uint64) ([]*rpcBlock, error) {
	var (
		elems  []rpc.BatchElem
		blocks = make([]*rpcBlock, len(blockNumbers))
	)

	for i, blockNumber := range blockNumbers {
		elems = append(elems, rpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args:   []interface{}{hexutil.EncodeUint64(blockNumber), false}, // 只取回区块里的交易hash值
			Result: &blocks[i],
		})
	}

	if err := rpcClient.BatchCall(elems); err != nil {
		return nil, errors.Wrap(err, "bat call elems failed")
	}

	return blocks, nil
}

func handleTxsInBlock(rpcClient *rpc.Client, block *rpcBlock) (resp *ScanLogsPtr, err error) {
	var (
		elems    []rpc.BatchElem
		receipts = make([]*types.Receipt, len(block.Transactions))
	)

	for i, tx := range block.Transactions {
		elems = append(elems, rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{tx.Hex()},
			Result: &receipts[i],
		})

		// 达到最大数量上限 提前处理一批
		if len(elems) >= maxReceipt {
			if err = rpcClient.BatchCall(elems); err != nil {
				return nil, errors.Wrapf(err, "batch get tx receipt failed, block number: %d", block.Number)
			}

			elems = elems[:0]
			continue
		}
	}

	if len(elems) > 0 {
		if err = rpcClient.BatchCall(elems); err != nil {
			return nil, errors.Wrapf(err, "batch get tx receipt failed, block number: %d", block.Number)
		}
	}

	logs := []*types.Log{}
	for _, receipt := range receipts {
		logs = append(logs, receipt.Logs...)
	}
	return &ScanLogsPtr{
		Logs:      logs,
		Timestamp: uint64(block.Timestamp),
	}, nil
}

func getTransactionReceipt(rpcClient *rpc.Client, txHash string) ([]*types.Log, error) {
	receipt := &types.Receipt{}

	if err := rpcClient.Call(receipt, "eth_getTransactionReceipt", common.HexToHash(txHash).Hex()); err != nil {
		return nil, errors.Wrapf(err, "batch get tx receipt failed, tx_hash: %s", txHash)
	}

	return receipt.Logs, nil
}
