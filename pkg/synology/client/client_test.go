package client

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/pwoelfle/synology-cli/testing/mock/server"
)

func TestClient_RequestGet_IfErrorResponse_ShouldReturnError(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.Return(`{"error":{"code":400},"success":false}`)

	sut := newClient(t, server.URL())

	// when:
	err := sut.requestGet("", map[string]string{}, nil)

	// then:
	assert.Equal(t, Error{Code: 400}, err)
}

func TestClient_RequestGet_ShouldLoadObjectAndReturnNil(t *testing.T) {
	// given:
	server := server.NewMockServer(t)
	defer server.Close()
	server.Return(`{"data":{"value":"test"},"success":true}`)

	sut := newClient(t, server.URL())

	// when:
	var object = &simpleObject{}
	err := sut.requestGet("", map[string]string{}, object)

	// then:
	assert.Nil(t, err)
	assert.Equal(t, "test", object.Value)
}

func newClient(t *testing.T, server string) *client {
	baseURL, err := url.Parse(server)
	require.Nil(t, err)

	return &client{
		baseURL:    *baseURL,
		httpClient: &http.Client{},
	}
}

var _ Object = new(simpleObject)

type simpleObject struct {
	Value string `json:"value"`
}
