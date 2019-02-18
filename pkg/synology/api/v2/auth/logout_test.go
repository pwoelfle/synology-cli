package auth

import (
	"testing"

	"github.com/pwoelfle/synology-cli/pkg/synology/api/v1/common"
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
	"github.com/pwoelfle/synology-cli/testing/mock/server"
	"github.com/stretchr/testify/assert"
)

func TestLogoutRequest_ShouldReturnNothing(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/auth.cgi",
		map[string]string{
			"api":     "SYNO.API.Auth",
			"version": "2",
			"method":  "logout",
			"session": "testing",
		}).
		Return(`{"success":true}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	req := NewLogoutRequest("testing")

	err = client.Call(req, nil)

	// then:
	assert.Nil(t, err)
}

func TestLogoutRequest_OnError_ShouldReturnNothing(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/auth.cgi",
		map[string]string{
			"api":     "SYNO.API.Auth",
			"version": "2",
			"method":  "logout",
			"session": "testing",
		}).
		Return(`{"error":{"code":100},"success":false}`)

	c, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	req := NewLogoutRequest("testing")

	err = c.Call(req, nil)

	// then:
	assert.NotNil(t, err)
	expectedError := client.Error{
		Code:    common.ErrorCodeUnknown,
		Message: "unknown error",
	}
	assert.Equal(t, expectedError, err)
}
