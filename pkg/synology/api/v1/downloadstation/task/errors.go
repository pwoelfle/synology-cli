package task

import (
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
)

const (
	ErrorCodeFileUploadFailed        client.ErrorCode = 400
	ErrorCodeMaxNumberOfTasksReached client.ErrorCode = 401
	ErrorCodeDestinationDenied       client.ErrorCode = 402
	ErrorCodeDestinationDoesNotExist client.ErrorCode = 403
	ErrorCodeInvalidTaskID           client.ErrorCode = 404
	ErrorCodeInvalidTaskAction       client.ErrorCode = 405
	ErrorCodeNoDefaultDestination    client.ErrorCode = 406
	ErrorCodeSetDestinationFailed    client.ErrorCode = 407
	ErrorCodeFileDoesNotExist        client.ErrorCode = 408
)
