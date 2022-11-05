package event

import (
	"dao-exchange/apps/fsm/dao"
	"dao-exchange/apps/indexer/contracts/nft"
	"dao-exchange/apps/indexer/parser/dex"

	log "github.com/sirupsen/logrus"
)

// OrderCancellAllEvent just impl actions interface
type OrderCancellAllEvent struct {
	Base
	parser *dex.MatchOrderParse
}

// NewOrderCancellAllEvent new obj
func NewOrderCancellAllEvent() *OrderCancellAllEvent {
	return &OrderCancellAllEvent{
		Base: Base{
			eventName:    OrderCancellAll,
			contractType: DAODex,
			eventDao:     &dao.BaseDaoOpr{},
		},
		parser: dex.NewMatchOrder(nft.NftMetaData.ABI),
	}
}

// HandlerMsg empty handler for reveiced msg
func (h *OrderCancellAllEvent) HandlerMsg(msg []byte) (interface{}, error) {
	log.Infof("OrderCancellAllEvent handling msg: %s", string(msg))
	return h.parser.Parse(msg)
}
