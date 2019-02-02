package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
	"github.com/pwoelfle/synology-cli/testing/mock/server"
)

func TestListRequest_ShouldReturnListOfTasks(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/DownloadStation/task.cgi",
		map[string]string{
			"api":     "SYNO.DownloadStation.Task",
			"version": "1",
			"method":  "list",
			"offset":  "0",
			"limit":   "-1",
		}).
		Return(`{"data":{"total":0,"offset":0,"tasks":[]},"success":true}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	taskList := &TaskList{}
	req := NewListRequest(ListRequestDefaultOffset, ListRequestDefaultLimit)

	err = client.Call(req, taskList)

	// then:
	assert.Nil(t, err)
	assert.Equal(t, 0, taskList.Total)
	assert.Equal(t, 0, taskList.Offset)
	assert.Equal(t, 0, len(taskList.Tasks))
}

func TestListRequest_WithAdditionals_ShouldReturnListOfTasks(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/DownloadStation/task.cgi",
		map[string]string{
			"api":        "SYNO.DownloadStation.Task",
			"version":    "1",
			"method":     "list",
			"offset":     "0",
			"limit":      "-1",
			"additional": "detail,file",
		}).
		Return(`{"data":{"total":0,"offset":0,"tasks":[]},"success":true}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	taskList := &TaskList{}
	req := NewListRequest(ListRequestDefaultOffset, ListRequestDefaultLimit, ListRequestAdditionalDetail, ListRequestAdditionalFile)

	err = client.Call(req, taskList)

	// then:
	assert.Nil(t, err)
	assert.Equal(t, 0, taskList.Total)
	assert.Equal(t, 0, taskList.Offset)
	assert.Equal(t, 0, len(taskList.Tasks))
}
