package main

import (
	_ "beego/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 如果无法访问，加入此配置即可。
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true

		// 开启 session, session 支持
		//beego.BConfig.WebConfig.Session.SessionOn = true

		// 内置了 admin 页面, 可以进行调试。
		//beego.BConfig.Listen.EnableAdmin = true
		//beego.BConfig.Listen.AdminAddr = "10.12.23.52"
		//beego.BConfig.Listen.AdminPort = 8888

		// 可开启内置, XSRF 机制。如果请求头中没有 对应key，则直接拒绝。
		//beego.BConfig.WebConfig.EnableXSRF = true
		//beego.BConfig.WebConfig.XSRFKey = "beego"
		//beego.BConfig.WebConfig.XSRFExpire = 3600

		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// beego 在运行的时候，内置了 init 方法，初始化了handler。

	// orm 的相关功能
	beego.Run()
}
