package task

import (
	"strconv"

	"github.com/pwoelfle/synology-cli/pkg/synology/client"
)

const (
	ListRequestDefaultOffset int = 0
	ListRequestDefaultLimit  int = -1
)

type ListRequestAdditional string

const (
	ListRequestAdditionalDetail   ListRequestAdditional = "detail"
	ListRequestAdditionalTransfer ListRequestAdditional = "transfer"
	ListRequestAdditionalFile     ListRequestAdditional = "file"
	ListRequestAdditionalTracker  ListRequestAdditional = "tracker"
	ListRequestAdditionalPeer     ListRequestAdditional = "peer"
)

func NewListRequest(offset, limit int, additionals ...ListRequestAdditional) client.Request {
	additionalString := ""

	for _, additional := range additionals {
		if len(additionalString) > 0 {
			additionalString = additionalString + ","
		}
		additionalString = additionalString + string(additional)
	}

	params := client.RequestParams{
		"offset": strconv.Itoa(offset),
		"limit":  strconv.Itoa(limit),
	}

	if len(additionalString) > 0 {
		params["additional"] = additionalString
	}

	return client.NewRequest(
		"DownloadStation/task.cgi",
		"SYNO.DownloadStation.Task",
		1,
		"list",
		params,
		taskErrorCodeMapper,
	)
}
