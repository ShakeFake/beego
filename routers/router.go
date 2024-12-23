// Package routers @APIVersion 1.0.0
// @Title plant_store
// @Description plant_store APIs
// @Contact
package routers

import (
	"beego/controllers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// autoRouter 下，使用全小写路径没错。
	// 如果同一个路由重复注册的话，会产生大量路由打印，这可以一眼看出来。
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoPrefix("api", &controllers.MainController{})

	// namespace 可以嵌套，没个嵌套下可以存在多个 ns。
	ns :=
		beego.NewNamespace("/api",
			// 进入 namespace 的必要条件。filter 的功能。注意，这个会影响所有功能。需要判空的。
			//beego.NSCond(func(c *context.Context) bool {
			//	return c.Request.Header["x-trace-id"][0] != ""
			//}),
			beego.NSNamespace("/Test",
				beego.NSInclude(
					&controllers.MainController{}, //user组对应的controller。用这里面的路径 + router。
				),
			),
		)

	// 另外一种方式，以filter方式过滤某些请求。
	//ns.Filter("before", func(ctx *context.Context) {
	//	fmt.Println("this is filter for health")
	//})
	// 注解路由

	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")

	//用来控制大小写不敏感的
	beego.WithCaseSensitive(true)

	// 注意下面这两种用法。这个涉及到是否适用指针问题。
	beego.CtrlGet("/api/user/:id", (*controllers.UserController).GetUserInfo)
	beego.CtrlGet("/api/user/no/:id", controllers.UserController.GetUserInfoNopointer)

	beego.CtrlPost("/api/user/:id", (*controllers.UserController).PostUserInfo)

	// 函数式注册方法。直接使用 function 丢入。
	beego.Get("/api/hello", controllers.Health)

	// 打印已经注册的方法。
	tree := beego.PrintTree()
	methods := tree["Data"].(beego.M)
	for k, v := range methods {
		fmt.Printf("%s => %v\n", k, v)
	}
}
