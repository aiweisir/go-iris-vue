package models

import (
	"go-iris/web/db"
	"time"

)

// mysql user table
type User struct {
	Id         int64  `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	Username   string `xorm:"notnull" json:"username" form:"username"`
	Password   string `xorm:"notnull" json:"password" form:"password"`
	Appid      string `xorm:"notnull" json:"appid" form:"appid"`
	Name       string `xorm:"notnull" json:"name" form:"name"`
	Phone      string `xorm:"notnull" json:"phone" form:"phone"`
	Email      string `xorm:"notnull" json:"email" form:"email"`
	Userface   string `xorm:"notnull" json:"userface" form:"userface"`
	CreateTime time.Time `json:"createTime" form:"createTime"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime"`
}

func CreateUser(user ...*User) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(user)
}

func GetUserByUsername(user *User) (bool, error) {
	e := db.MasterEngine()
	return e.Get(user)
}
