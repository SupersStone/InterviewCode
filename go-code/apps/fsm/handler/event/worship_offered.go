package event

import (
	log "github.com/sirupsen/logrus"

	"dao-exchange/apps/fsm/dao"
	contract_worship "dao-exchange/apps/indexer/contracts/nft/worship"
	"dao-exchange/apps/indexer/parser/worship"
)

// WorshipOfferedEvent just impl actions interface
type WorshipOfferedEvent struct {
	Base
	parser *worship.OfferedParse
}

// NewWorshipOfferedEvent new obj
func NewWorshipOfferedEvent() *WorshipOfferedEvent {
	return &WorshipOfferedEvent{
		Base: Base{
			eventName:    WorshipOffered,
			contractType: Worship,
			eventDao:     &dao.BaseDaoOpr{},
		},
		parser: worship.NewOfferedParse(contract_worship.NftMetaData.ABI),
	}
}

// HandlerMsg empty handler for reveiced msg
func (h *WorshipOfferedEvent) HandlerMsg(msg []byte) (interface{}, error) {
	log.Infof("WorshipOfferedEvent handling msg: %s", string(msg))
	return h.parser.Parse(msg)
}
