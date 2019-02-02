package client

import "fmt"

var _ error = new(Error)

type ErrorCode int

type Error struct {
	Code ErrorCode `json:"code"`
}

func (e Error) Error() string {
	return fmt.Sprintf("error code %d", e.Code)
}
