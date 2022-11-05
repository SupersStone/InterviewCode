package mq

import (
	"github.com/ethereum/go-ethereum/common"
)

// SendMsg send message
type SendMsg struct {
	ContractAddr common.Address
	Data         []byte
}
