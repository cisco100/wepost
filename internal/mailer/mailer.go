package mailer

import "embed"

const (
	FromName            string = "WePost"
	MaxRetriesLimit     int    = 3
	UserInvitesTEmplate string = "user_invitations.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) error
}
