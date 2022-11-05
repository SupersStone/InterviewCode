package event

import (
	log "github.com/sirupsen/logrus"

	"dao-exchange/apps/fsm/dao"
	nftToken "dao-exchange/apps/indexer/contracts/nft/token"
	"dao-exchange/apps/indexer/parser/erc721"
)

// TransferEvent just impl actions interface
type TransferEvent struct {
	Base
	parser *erc721.TransferParse
}

// NewTransferEvent new obj
func NewTransferEvent() *TransferEvent {
	return &TransferEvent{
		Base: Base{
			eventName:    TransferToken,
			contractType: Erc721,
			eventDao:     &dao.BaseDaoOpr{},
		},
		parser: erc721.NewTransferParse(nftToken.TokenMetaData.ABI),
	}
}

// HandlerMsg empty handler for reveiced msg
func (h *TransferEvent) HandlerMsg(msg []byte) (interface{}, error) {
	log.Infof("TransferEvent handling msg: %s", string(msg))
	return h.parser.Parse(msg)
}
