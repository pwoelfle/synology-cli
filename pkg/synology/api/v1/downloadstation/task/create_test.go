package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
	"github.com/pwoelfle/synology-cli/testing/mock/server"
)

func TestCreateRequest_IfSuccess_ShouldNothing(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/DownloadStation/task.cgi",
		map[string]string{
			"api":     "SYNO.DownloadStation.Task",
			"version": "1",
			"method":  "create",
			"uri":     "https://www.example.com/files/download.php?id=1234",
		}).
		Return(`{"success":true}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	req := NewCreateRequest("https://www.example.com/files/download.php?id=1234")

	err = client.Call(req, nil)

	// then:
	assert.Nil(t, err)
}

func TestCreateRequest_IfNotSuccess_ShouldReturnError(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/DownloadStation/task.cgi",
		map[string]string{
			"api":     "SYNO.DownloadStation.Task",
			"version": "1",
			"method":  "create",
			"uri":     "https://www.example.com/files/download.php?id=1234",
		}).
		Return(`{"error":{"code":123},"success":false}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	req := NewCreateRequest("https://www.example.com/files/download.php?id=1234")

	err = client.Call(req, nil)

	// then:
	assert.NotNil(t, err)
}
