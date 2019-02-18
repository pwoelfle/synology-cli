package task

import (
	"github.com/pwoelfle/synology-cli/pkg/synology/api/v1/common"
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

func taskErrorCodeMapper(errCode client.ErrorCode) string {
	switch errCode {
	case ErrorCodeFileUploadFailed:
		return "file upload failed"
	case ErrorCodeMaxNumberOfTasksReached:
		return "max number of tasks reached"
	case ErrorCodeDestinationDenied:
		return "destination denied"
	case ErrorCodeDestinationDoesNotExist:
		return "destination does not exist"
	case ErrorCodeInvalidTaskID:
		return "invalid task id"
	case ErrorCodeInvalidTaskAction:
		return "invalid task action"
	case ErrorCodeNoDefaultDestination:
		return "no default destination"
	case ErrorCodeSetDestinationFailed:
		return "set destination failed"
	case ErrorCodeFileDoesNotExist:
		return "file does not exist"
	}
	return common.CommonErrorCodeMapper(errCode)
}
