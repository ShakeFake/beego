package controllers

import (
	"beego/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"time"
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
	name := u.GetString("name", "zhangsan", "zhangsan_moren")
	current := time.Now()
	toys := []*models.Toy{
		{
			ID:     1,
			Name:   "abc",
			UserID: 1,
		},
		{
			ID:     2,
			Name:   "abc_2",
			UserID: 1,
		},
	}

	cc := models.CreditCard{
		ID:     1,
		Number: "1",
		UserID: 1,
	}

	// 写在一起，并不能同时追加进去。如果要保证关系存在，需要显示设置一下。
	user := models.User{
		ID:       1,
		Name:     name,
		Birthday: &current,
	}

	fmt.Println(user)

	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()

	// 这个地方，需要使用指针。 beego 需要一条一条的插入。
	id, err := o.Insert(&user)
	fmt.Println(id, err)
	id, err = o.InsertMulti(3, &toys)
	fmt.Println(id, err)
	id, err = o.Insert(&cc)
	fmt.Println(id, err)

	m2m := o.QueryM2M(&user, "Toys")
	num, err := m2m.Add(toys)
	fmt.Println(num, err)

	u.Ctx.WriteString("insert person success")
}

func (u *UserController) PutUserInfo() {

	u.Ctx.WriteString("put user info")
}

func (u *UserController) DeleteUserInfo() {
	u.Ctx.WriteString("delete user info")
}

func (u *UserController) PostFile() {
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
