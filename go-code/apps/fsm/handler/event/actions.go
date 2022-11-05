package event

import "gorm.io/gorm"

// Actions consumer msg and handler it with actions
type Actions interface {
	HandlerMsg(msg []byte) (interface{}, error)
	InsertDB(db *gorm.DB, data interface{}) error
}

// ContractType contract type alias
type ContractType string

// contract type
const (
	DAODex  ContractType = "dao-dex"
	Erc721  ContractType = "erc721"
	Proxy   ContractType = "proxy"
	Worship ContractType = "wroship"
)

const (
	// OrderCancelled the event name
	OrderCancelled = "OrderCancelled"
	// OrderCancellAll the event name
	OrderCancellAll = "AllOrdersCancelled"
	// FixedPriceOrderMatched the event name
	FixedPriceOrderMatched = "FixedPriceOrderMatched"
	// TransferToken the event name
	TransferToken = "Transfer"
	// WorshipOffered the event name
	WorshipOffered = "WorshipOffered"
	// WorshipRedeemed the event name
	WorshipRedeemed = "WorshipRedeemed"
)
