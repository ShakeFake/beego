package main

import (
	_ "beego/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 如果无法访问，加入此配置即可。
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
