package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wung-s/ukhrul/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
