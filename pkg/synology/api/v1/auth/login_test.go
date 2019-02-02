package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/pwoelfle/synology-cli/pkg/synology/client"
	"github.com/pwoelfle/synology-cli/testing/mock/server"
)

func TestLoginRequest_ShouldReturnSessionID(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.OnPathWithParameter(
		"/webapi/auth.cgi",
		map[string]string{
			"api":     "SYNO.API.Auth",
			"version": "1",
			"method":  "login",
			"account": "user",
			"passwd":  "pwd",
			"session": "testing",
		}).
		Return(`{"data":{"sid":"abcdef"},"success":true}`)

	client, err := client.NewClient(server.URL())
	assert.Nil(t, err)

	// when:
	session := &Session{}
	req := NewLoginRequest("user", "pwd", "testing")

	err = client.Call(req, session)

	// then:
	assert.Nil(t, err)
	assert.Equal(t, "abcdef", session.SessionID)
}
