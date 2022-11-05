package myerr

import "errors"

// define err
var (
	ErrRecordNotFound = errors.New("record not found")
)
