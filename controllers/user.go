package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (u *UserController) HelloWorld() {
	u.Ctx.WriteString("hello, world")
}

type User struct {
	Name string `json:"name" form:"name"`
}

func (u *UserController) GetUserInfo() {
	// 参数 > 表单。虽然 GetString 接收多个参数，但是只会使用第一个值。
	name := u.GetString("name", "zhangsan")

	user := User{}
	if err := u.BindForm(&user); err != nil {
		u.Ctx.WriteString(fmt.Sprintf("bind form error: %v", err))
		return
	}

	// cookie 支持
	//  u.Ctx.GetCookie()
	//  u.Ctx.SetCookie()
	//u.Ctx.SetSecureCookie()
	//u.Ctx.GetSecureCookie()

	u.Ctx.WriteString(fmt.Sprintf("get user info: %v_%v", name, user.Name))
}

func (u UserController) GetUserInfoNopointer() {
	u.Ctx.WriteString("get user info no pointer")
}

func (u *UserController) PostUserInfo() {
	// 注意请求头和存储名字。
	f, h, err := u.GetFile("uploadname")
	if err != nil {
		u.Ctx.WriteString(fmt.Sprintf("get file error: %v", err))
		return
	}
	defer f.Close()

	// filename 带格式的
	u.SaveToFile("uploadname", "./"+h.Filename)
	u.Ctx.WriteString("save file success")
}

func (u *UserController) PutUserInfo() {
	u.Ctx.WriteString("put user info")
}

func (u *UserController) DeleteUserInfo() {
	u.Ctx.WriteString("delete user info")
}
