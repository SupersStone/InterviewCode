package event

import (
	"dao-exchange/apps/fsm/dao"
	nftToken "dao-exchange/apps/indexer/contracts/nft/token"
	"dao-exchange/apps/indexer/parser/erc721"

	log "github.com/sirupsen/logrus"
)

/*
	TODO 需要考虑一个问题，创建proxy 的时候，地址可以重复，但是proxy不可以重复，
	所以这里需要考虑的问题是，每次是更新还是直接插入新的
*/

// CreateProxyEvent just impl actions interface
type CreateProxyEvent struct {
	Base
	parser *erc721.TransferParse
}

// NewCreateProxyEvent new obj
func NewCreateProxyEvent() *CreateProxyEvent {
	return &CreateProxyEvent{
		Base: Base{
			eventName:    TransferToken,
			contractType: Erc721,
			eventDao:     &dao.BaseDaoOpr{},
		},
		parser: erc721.NewTransferParse(nftToken.TokenMetaData.ABI),
	}
}

// HandlerMsg empty handler for reveiced msg
func (h *CreateProxyEvent) HandlerMsg(msg []byte) (interface{}, error) {
	log.Infof("CreateProxyEvent handling msg: %s", string(msg))
	return h.parser.Parse(msg)
}
