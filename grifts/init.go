package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wung-s/buffalo-auth0-authentication/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
