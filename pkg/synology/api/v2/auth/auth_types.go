package auth

type LoginFormat string

const (
	LoginFormatCookie LoginFormat = "cookie"
	LoginFormatSid    LoginFormat = "sid"
)

type Session struct {
	SessionID string `json:"sid"`
}
