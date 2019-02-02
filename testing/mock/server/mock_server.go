package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockServer interface {
	URL() string

	OnPath(path string) MockServer

	OnPathWithParameter(path string, params map[string]string) MockServer

	Return(response string) MockServer

	ReturnError() MockServer

	Close()
}

type mockServer struct {
	t      *testing.T
	server *httptest.Server

	path   string
	params map[string]string

	responseError bool
	response      string
}

func (m *mockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	assert.Equal(m.t, m.path, r.URL.Path, "request path '%s' does not match with mock path '%s'", r.URL.Path, m.path)

	requestParams := r.URL.Query()
	assert.Equal(m.t, len(m.params), len(requestParams), "number of request params (%d) does not match with number of mock params (%d)", len(requestParams), len(m.params))

	for k, v := range m.params {
		requestParamValue := requestParams.Get(k)
		assert.Equal(m.t, v, requestParamValue, "for parameter key '%s', request param value '%s' does not match mock param value '%s'", k, requestParamValue, v)
	}

	if m.responseError {
		http.Error(w, "an error occured", 505)
	} else {
		fmt.Fprint(w, m.response)
	}
}

func (m *mockServer) URL() string {
	return m.server.URL
}

func (m *mockServer) OnPath(path string) MockServer {
	m.path = path

	return m
}

func (m *mockServer) OnPathWithParameter(path string, params map[string]string) MockServer {
	m.path = path
	m.params = params

	return m
}

func (m *mockServer) Return(response string) MockServer {
	m.responseError = false
	m.response = response

	return m
}

func (m *mockServer) ReturnError() MockServer {
	m.responseError = true
	m.response = ""

	return m
}

func (m *mockServer) Close() {
	m.server.Close()
}

func NewMockServer(t *testing.T) MockServer {
	mockServer := &mockServer{
		t: t,

		path:   "/",
		params: map[string]string{},

		response: "{}",
	}

	mockServer.server = httptest.NewServer(mockServer)
	return mockServer
}
