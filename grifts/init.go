package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/user/samplebuffalo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
