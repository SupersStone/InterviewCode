package mocks

import "fmt"

// EmptyParser just impl actions interface
type EmptyParser struct {
}

// HandlerMsg empty handler for reveiced msg
func (p *EmptyParser) HandlerMsg([]byte) error {
	fmt.Println("handling empty msg")
	return nil
}
