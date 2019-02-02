package auth

import (
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
)

const (
	ErrorCodeAuthFailed        client.ErrorCode = 400
	ErrorCodeAccountDisabled   client.ErrorCode = 401
	ErrorCodePermissionDenied  client.ErrorCode = 402
	ErrorCodeTwoFactorRequired client.ErrorCode = 403
	ErrorCodeTwoFactorFailed   client.ErrorCode = 404
)
