package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type MainController struct {
	beego.Controller
}

// @Title CheckHelath
// @Description just health check
// @Success 201 success
// @Failure 500 failure
// @router /health
func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
	return
}

func Health(ctx *context.Context) {
	ctx.WriteString("Success")
	return
}
