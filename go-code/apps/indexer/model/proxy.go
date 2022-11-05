package model

import (
	"github.com/ethereum/go-ethereum/common"
)

// CreateProxyEvent event
type CreateProxyEvent struct {
	User    common.Address
	Conduit common.Address
}
