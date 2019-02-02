package auth

import "github.com/pwoelfle/synology-cli/pkg/synology/client"

func NewLoginRequest(username, password, session string, format LoginFormat) client.Request {
	return client.NewRequest(
		"auth.cgi",
		"SYNO.API.Auth",
		2,
		"login",
		client.RequestParams{
			"account": username,
			"passwd":  password,
			"session": session,
			"format":  string(format),
		},
	)
}
