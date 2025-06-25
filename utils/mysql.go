package utils

import (
	"beego/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"

	// 数据库驱动需要匿名引入
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.CreditCard))
	orm.RegisterModel(new(models.Toy))
	//orm.RegisterModelWithPrefix("tab_", &models.User{})
	//orm.RegisterModelWithSuffix("_tab", &models.User{})

	dsn := "root:root@tcp(192.168.1.1:3306)/beego?charset=utf8mb4&parseTime=True&loc=Local"
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.SetMaxOpenConns("default", 30)
	orm.SetMaxIdleConns("default", 20)

	// 用来设置默认数据库时间作用
	orm.DefaultTimeLoc = time.UTC

	fmt.Println("register success")
}
