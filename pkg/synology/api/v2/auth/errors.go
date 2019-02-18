package auth

import (
	"github.com/pwoelfle/synology-cli/pkg/synology/api/v2/common"
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
)

const (
	ErrorCodeAuthFailed        client.ErrorCode = 400
	ErrorCodeAccountDisabled   client.ErrorCode = 401
	ErrorCodePermissionDenied  client.ErrorCode = 402
	ErrorCodeTwoFactorRequired client.ErrorCode = 403
	ErrorCodeTwoFactorFailed   client.ErrorCode = 404
)

func authErrorCodeMapper(errCode client.ErrorCode) string {
	switch errCode {
	case ErrorCodeAuthFailed:
		return "authentication failed"
	case ErrorCodeAccountDisabled:
		return "account disabled"
	case ErrorCodePermissionDenied:
		return "permission denied"
	case ErrorCodeTwoFactorRequired:
		return "two-factor authentication required"
	case ErrorCodeTwoFactorFailed:
		return "two-factor authentication failed"
	}
	return common.CommonErrorCodeMapper(errCode)
}
