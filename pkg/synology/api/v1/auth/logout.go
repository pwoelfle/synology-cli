package auth

import "github.com/pwoelfle/synology-cli/pkg/synology/client"

func NewLogoutRequest(session string) client.Request {
	return client.NewRequest(
		"auth.cgi",
		"SYNO.API.Auth",
		1,
		"logout",
		client.RequestParams{
			"session": session,
		},
		authErrorCodeMapper,
	)
}
