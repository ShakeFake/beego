package models

import (
	"time"
)

// User beego 创造出来的 User 表全是非空字段。
type User struct {
	// auto: 自增主键。 pk 主键。
	ID   int    `orm:"column(id);index;unique"`
	Name string `orm:"column(name)"`

	// auto_now_add 第一次添加时更新。 auto_now 每次保存时均更新。
	Birthday *time.Time `orm:"auto_now_add;type(datetime)"`

	// reverse 用来支持反向关系。one 一对一。 fk外键。many一对多。m2m多对多关系。
	// 多对多关系额外建表。一对一关系，存储在主表。本例即 user 表。
	// on_delete支持字段。 cascade级联删除；set_null；set_default；do_nothing；
	CreditCard *CreditCard `orm:"null;rel(one);on_delete(set_null)"`
	Toys       []*Toy      `orm:"null;rel(m2m);on_delete(do_nothing)"`
}

// TableName 自定义表名，实现
func (u *User) TableName() string {
	return "user"
}

// TableIndex 设置多字段联合索引
func (u *User) TableIndex() [][]string {
	return [][]string{
		{"ID", "Name"},
	}
}

// TableUnique 设置多字段唯一键
func (u *User) TableUnique() [][]string {
	return [][]string{
		{"Name"},
	}
}

type CreditCard struct {
	ID     int
	Number string
	UserID uint
}

type Toy struct {
	ID   int
	Name string
	Kind string

	UserID   int
	UserType string
}
