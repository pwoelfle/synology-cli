package auth

import "github.com/pwoelfle/synology-cli/pkg/synology/client"

func NewLoginRequest(username, password, session string) client.Request {
	return client.NewRequest(
		"auth.cgi",
		"SYNO.API.Auth",
		1,
		"login",
		client.RequestParams{
			"account": username,
			"passwd":  password,
			"session": session,
		},
	)
}
