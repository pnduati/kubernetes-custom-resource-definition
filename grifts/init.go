package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pnduati/kubernetes_custom_resource_definition/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
