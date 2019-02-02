package common

import (
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
)

const (
	ErrorCodeUnknown                  client.ErrorCode = 100
	ErrorCodeInvalidParameter         client.ErrorCode = 101
	ErrorCodeAPINotExists             client.ErrorCode = 102
	ErrorCodeMethodNotExists          client.ErrorCode = 103
	ErrorCodeUnsupportedFunctionality client.ErrorCode = 104
	ErrorCodeMissingPermission        client.ErrorCode = 105
	ErrorCodeSessionTimeout           client.ErrorCode = 106
	ErrorCodeDuplicateLogin           client.ErrorCode = 107
)
