package task

import (
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
)

func NewCreateRequest(uri string) client.Request {
	params := client.RequestParams{
		"uri": uri,
	}

	return client.NewRequest(
		"DownloadStation/task.cgi",
		"SYNO.DownloadStation.Task",
		1,
		"create",
		params,
		taskErrorCodeMapper,
	)
}
