package vo

import "time"

// 前端需要的数据结构
type UserVO struct {
	Id         int64     `json:"id" form:"id"`
	Username   string    `json:"username" form:"username"`
	Appid      string    `json:"appid" form:"appid"`
	Name       string    `json:"name" form:"name"`
	Phone      string    `json:"phone" form:"phone"`
	Email      string    `json:"email" form:"email"`
	Userface   string    `json:"userface" form:"userface"`
	CreateTime time.Time `json:"createTime" form:"createTime"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime"`
	Token      string    `json:"token"`
}
