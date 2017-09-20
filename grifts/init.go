package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/buffla/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
