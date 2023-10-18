// Package routers @APIVersion 1.0.0
// @Title plant_store
// @Description plant_store APIs
// @Contact
package routers

import (
	"beego/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns :=
		beego.NewNamespace("/api",
			beego.NSNamespace("/test",
				beego.NSInclude(
					&controllers.MainController{}, //user组对应的controller
				),
			),
		)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")
	beego.Router("/api/test/health", &controllers.MainController{})
}
