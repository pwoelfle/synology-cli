package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
	"github.com/pwoelfle/synology-cli/testing/mock/server"
)

func TestGetInfoRequest_ShouldReturnListOfTasks(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/DownloadStation/task.cgi",
		map[string]string{
			"api":     "SYNO.DownloadStation.Task",
			"version": "1",
			"method":  "getinfo",
			"id":      "0",
		}).
		Return(`{"data":{"tasks":[]},"success":true}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	taskGetInfo := &TaskGetInfo{}
	req := NewGetInfoRequest([]ID{"0"})

	err = client.Call(req, taskGetInfo)

	// then:
	assert.Nil(t, err)
	assert.Equal(t, 0, len(taskGetInfo.Tasks))
}

func TestGetInfoRequest_WithMultipleIds_ShouldReturnListOfTasks(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/DownloadStation/task.cgi",
		map[string]string{
			"api":     "SYNO.DownloadStation.Task",
			"version": "1",
			"method":  "getinfo",
			"id":      "0,1,2",
		}).
		Return(`{"data":{"tasks":[]},"success":true}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	taskGetInfo := &TaskGetInfo{}
	req := NewGetInfoRequest([]ID{"0", "1", "2"})

	err = client.Call(req, taskGetInfo)

	// then:
	assert.Nil(t, err)
	assert.Equal(t, 0, len(taskGetInfo.Tasks))
}

func TestGetInfoRequest_WithAdditionals_ShouldReturnListOfTasks(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/DownloadStation/task.cgi",
		map[string]string{
			"api":        "SYNO.DownloadStation.Task",
			"version":    "1",
			"method":     "getinfo",
			"id":         "0",
			"additional": "detail,file",
		}).
		Return(`{"data":{"tasks":[]},"success":true}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	taskGetInfo := &TaskGetInfo{}
	req := NewGetInfoRequest([]ID{"0"}, GetInfoRequestAdditionalDetail, GetInfoRequestAdditionalFile)

	err = client.Call(req, taskGetInfo)

	// then:
	assert.Nil(t, err)
	assert.Equal(t, 0, len(taskGetInfo.Tasks))
}
