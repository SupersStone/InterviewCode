package event

import (
	"dao-exchange/apps/fsm/dao"
	"dao-exchange/apps/indexer/contracts/nft"
	"dao-exchange/apps/indexer/parser/dex"

	log "github.com/sirupsen/logrus"
)

// MatchEvent just impl actions interface
type MatchEvent struct {
	Base
	parser *dex.MatchOrderParse
}

// NewMatchEvent new obj
func NewMatchEvent() *MatchEvent {
	return &MatchEvent{
		Base: Base{
			eventName:    FixedPriceOrderMatched,
			contractType: DAODex,
			eventDao:     &dao.BaseDaoOpr{},
		},
		parser: dex.NewMatchOrder(nft.NftMetaData.ABI),
	}
}

// HandlerMsg empty handler for reveiced msg
func (h *MatchEvent) HandlerMsg(msg []byte) (interface{}, error) {
	log.Infof("MatchEvent handling msg: %s", string(msg))
	return h.parser.Parse(msg)
}
