package evms

import "github.com/ethereum/go-ethereum/core/types"

// ScanLogsPtr scan response
type ScanLogsPtr struct {
	Logs      []*types.Log
	Timestamp uint64
}

// ScanLogs scan response
type ScanLogs struct {
	Logs      []types.Log
	Timestamp uint64
}
