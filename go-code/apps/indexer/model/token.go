package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TokenTransferEvent event
type TokenTransferEvent struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}
