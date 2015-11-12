package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	GorpController
}

func (c *App) Before() revel.Result {
	// Rendering useful info here.
	c.RenderArgs["action"] = c.Controller.Action

	return nil
}

func init() {
	revel.InterceptMethod((*App).Before, revel.BEFORE)
}
