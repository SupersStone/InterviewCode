package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// OrderMatchEvent construct event info, because event was nested structure
type OrderMatchEvent struct {
	OriMatchEvent FixedPriceOrderMatch
	OriOrderInfo  OrderBytesInfo
	OriTokenInfo  TokenBytesInfo
}

// FixedPriceOrderMatch event
type FixedPriceOrderMatch struct {
	Maker       common.Address `json:"maker"`
	Taker       common.Address `json:"taker"`
	OrderHash   common.Hash    `json:"order_digest"`
	OrderBytes  []byte         `json:"order_bytes"`
	AssetsBytes []byte         `json:"assets_bytes"`
}

// OrderBytesInfo for resolve order struct
type OrderBytesInfo struct {
	OrderStructHash  common.Hash    `json:"order_struct_hash"`
	Maker            common.Address `json:"maker"`
	Taker            common.Address `json:"taker"`
	RoyaltyRecipient common.Address `json:"royalty_recipient"`
	RoyaltyRate      *big.Int       `json:"royalty_rate"`
	StartAt          uint64         `json:"start_at"`
	ExpireAt         uint64         `json:"expire_at"`
	MakerNonce       uint64         `json:"make_nonce"`
	TakerGetNft      bool           `json:"taker_get_nft"`
	AssetsHash       common.Hash    `json:"assets_hash"`
}

// TokenBytesInfo for resolve order token struct
type TokenBytesInfo struct {
	AssetsStructHash common.Hash    `json:"assets_struct_hash"`
	Nft              common.Address `json:"nft"`
	Ft               common.Address `json:"ft"`
	NftId            *big.Int       `json:"nft_id"`
	NftAmount        *big.Int       `json:"nft_amount"`
	FtAmount         *big.Int       `json:"ft_amount"`
}
