package client

import "fmt"

var _ error = new(Error)

type ErrorCode int

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}
