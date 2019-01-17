package vo

import (
	"go-iris/web/models"

	"github.com/kataras/golog"
)

// 前端需要的数据结构
type UserVO struct {
	Id         int64  `json:"id" form:"id"`
	Username   string `json:"username" form:"username"`
	Enable     int    `json:"enable"`
	Appid      string `json:"appid" form:"appid"`
	Name       string `json:"name" form:"name"`
	Phone      string `json:"phone" form:"phone"`
	Email      string `json:"email" form:"email"`
	Userface   string `json:"userface" form:"userface"`
	CreateTime int64  `json:"create_time" form:"createTime"`
	UpdateTime int64  `json:"update_time" form:"updateTime"`
	Token      string `json:"token"`
}

// 携带token
func TansformUserVO(token string, user *models.User) (uVO UserVO) {
	uVO.Id = user.Id
	uVO.Username = user.Username
	uVO.Enable = user.Enable
	uVO.Appid = user.Appid
	uVO.Name = user.Name
	uVO.Phone = user.Phone
	uVO.Email = user.Email
	uVO.Userface = user.Userface
	uVO.CreateTime = user.CreateTime.Unix()
	uVO.UpdateTime = user.UpdateTime.Unix()

	uVO.Token = token

	golog.Infof("uptime=%d", user.UpdateTime.UnixNano())
	return
}

// 用户列表，不带啊token
func TansformUserVOList(userList ...*models.User) (userVOList []UserVO) {
	for _, v := range userList {
		uVO := UserVO{}
		uVO.Id = v.Id
		uVO.Username = v.Username
		uVO.Enable = v.Enable
		uVO.Appid = v.Appid
		uVO.Name = v.Name
		uVO.Phone = v.Phone
		uVO.Email = v.Email
		uVO.Userface = v.Userface
		uVO.CreateTime = v.CreateTime.Unix()
		uVO.UpdateTime = v.UpdateTime.Unix()

		userVOList = append(userVOList, uVO)
	}
	return
}
