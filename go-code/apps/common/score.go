package common

import "math/big"

func CalcScore(blockNum, txIdx, logIdx uint64) string {
	blockNumber := big.NewInt(0).SetUint64(blockNum)
	txIndex := big.NewInt(0).SetUint64(txIdx)
	logIndex := big.NewInt(0).SetUint64(logIdx)
	blockNumber = big.NewInt(0).Mul(blockNumber, big.NewInt(100000000))
	txIndex = big.NewInt(0).Mul(txIndex, big.NewInt(10000))
	score := big.NewInt(0).Add(logIndex, big.NewInt(0).Add(blockNumber, txIndex))
	return score.String()
}
