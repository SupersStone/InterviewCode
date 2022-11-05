package nft

import "github.com/ethereum/go-ethereum/crypto"

var (
	logOrderCancelledSig         = []byte("OrderCancelled(maker,order_digest)")
	logAllOrdersCancelledSig     = []byte("AllOrdersCancelled(maker,current_nonce)")
	logFixedPriceOrderMatchedSig = []byte("FixedPriceOrderMatched(sender,maker,taker,maker_order_digest,taker_order_digest,maker_order_bytes,taker_order_bytes,tokens_bytes)")
	// LogOrderCancelledSigHash transfer func hash
	LogOrderCancelledSigHash = crypto.Keccak256Hash(logOrderCancelledSig)
	// LogAllOrdersCancelledSigHash approval func hash
	LogAllOrdersCancelledSigHash = crypto.Keccak256Hash(logAllOrdersCancelledSig)
	// LogFixedPriceOrderMatchedSigHash approval func hash
	LogFixedPriceOrderMatchedSigHash = crypto.Keccak256Hash(logFixedPriceOrderMatchedSig)
)
