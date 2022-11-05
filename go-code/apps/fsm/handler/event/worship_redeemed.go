package event

import (
	log "github.com/sirupsen/logrus"

	"dao-exchange/apps/fsm/dao"
	contract_worship "dao-exchange/apps/indexer/contracts/nft/worship"
	"dao-exchange/apps/indexer/parser/worship"
)

// WorshipRedeemedEvent just impl actions interface
type WorshipRedeemedEvent struct {
	Base
	parser *worship.RedeemedParse
}

// NewWorshipRedeemedEvent new obj
func NewWorshipRedeemedEvent() *WorshipRedeemedEvent {
	return &WorshipRedeemedEvent{
		Base: Base{
			eventName:    WorshipRedeemed,
			contractType: Worship,
			eventDao:     &dao.BaseDaoOpr{},
		},
		parser: worship.NewRedeemedParse(contract_worship.NftMetaData.ABI),
	}
}

// HandlerMsg empty handler for reveiced msg
func (h *WorshipRedeemedEvent) HandlerMsg(msg []byte) (interface{}, error) {
	log.Infof("WorshipRedeemedEvent handling msg: %s", string(msg))
	return h.parser.Parse(msg)
}
