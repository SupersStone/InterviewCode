package event

import (
	"dao-exchange/apps/fsm/dao"
	"dao-exchange/apps/indexer/contracts/nft"
	"dao-exchange/apps/indexer/parser/dex"

	log "github.com/sirupsen/logrus"
)

// OrderCancelledEvent just impl actions interface
type OrderCancelledEvent struct {
	Base
	parser *dex.MatchOrderParse
}

// NewOrderCancelledEvent new obj
func NewOrderCancelledEvent() *OrderCancelledEvent {
	return &OrderCancelledEvent{
		Base: Base{
			eventName:    OrderCancelled,
			contractType: DAODex,
			eventDao:     &dao.BaseDaoOpr{},
		},
		parser: dex.NewMatchOrder(nft.NftMetaData.ABI),
	}
}

// HandlerMsg empty handler for reveiced msg
func (h *OrderCancelledEvent) HandlerMsg(msg []byte) (interface{}, error) {
	log.Infof("OrderCancelledEvent handling msg: %s", string(msg))
	return h.parser.Parse(msg)
}
