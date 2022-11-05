package evms

import "github.com/ethereum/go-ethereum/core/types"

// Msg message for kafka
type Msg struct {
	Log            *types.Log `json:"log"`
	ChainID        int        `json:"chainId"`
	ChainName      string     `json:"chainName"`
	BlockTimestamp uint64     `json:"blockTimestamp"`
}
