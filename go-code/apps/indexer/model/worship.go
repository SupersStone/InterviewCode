package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TokenOfferedEvent event
type TokenOfferedEvent struct {
	Nft              common.Address `json:"nft"`
	TokenId          *big.Int       `json:"tokenId"`
	Votary           common.Address `json:"votary"`
	ReleaseTimestamp *big.Int       `json:"releaseTimestamp"`
	Redeemer         common.Address `json:"redeemer"`
}

// TokenRedeemedEvent for resolve order struct
type TokenRedeemedEvent struct {
	Nft      common.Address `json:"nft"`
	TokenId  *big.Int       `json:"tokenId"`
	Redeemer common.Address `json:"redeemer"`
}
