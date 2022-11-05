package common

import "errors"

// errors definiton
var (
	ErrNotFoundSellOrder  = errors.New("not found sell order")
	ErrNotFoundEventDef   = errors.New("not found event definiton")
	ErrUnMarshalFail      = errors.New("unmarshal err")
	ErrNotSupportEvent    = errors.New("not support event")
	ErrSystemMetaNotFound = errors.New("system metadata not found")
	ErrDexNotMatchEvent   = errors.New("dex not match with event")
	ErrNotSupportChainID  = errors.New("not support chain id")
	ErrNotSupportPayToken = errors.New("not support pay token")
)
