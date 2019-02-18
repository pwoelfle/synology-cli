package task

import (
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
)

type GetInfoRequestAdditional string

const (
	GetInfoRequestAdditionalDetail   GetInfoRequestAdditional = "detail"
	GetInfoRequestAdditionalTransfer GetInfoRequestAdditional = "transfer"
	GetInfoRequestAdditionalFile     GetInfoRequestAdditional = "file"
	GetInfoRequestAdditionalTracker  GetInfoRequestAdditional = "tracker"
	GetInfoRequestAdditionalPeer     GetInfoRequestAdditional = "peer"
)

func NewGetInfoRequest(ids []ID, additionals ...GetInfoRequestAdditional) client.Request {
	idString := ""

	for _, id := range ids {
		if len(idString) > 0 {
			idString = idString + ","
		}
		idString = idString + string(id)
	}

	additionalString := ""

	for _, additional := range additionals {
		if len(additionalString) > 0 {
			additionalString = additionalString + ","
		}
		additionalString = additionalString + string(additional)
	}

	params := client.RequestParams{
		"id": idString,
	}

	if len(additionalString) > 0 {
		params["additional"] = additionalString
	}

	return client.NewRequest(
		"DownloadStation/task.cgi",
		"SYNO.DownloadStation.Task",
		1,
		"getinfo",
		params,
		taskErrorCodeMapper,
	)
}
