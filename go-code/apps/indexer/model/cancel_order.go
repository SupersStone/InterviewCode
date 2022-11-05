package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CancelOrderEvent construct event info, because event was nested structure
type CancelOrderEvent struct {
	Maker     common.Address `json:"maker"`
	OrderHash common.Hash    `json:"order_hash"`
}

// CancelAllOrderEvent construct event info, because event was nested structure
type CancelAllOrderEvent struct {
	Offerer        common.Address `json:"offerer"`
	IncreasedNonce *big.Int       `json:"increased_nonce"`
}
