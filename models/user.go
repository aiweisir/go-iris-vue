package models

import "time"

type User struct {
	Id int `gorm:"primary_key"`
	Username string `gorm:"type:varchar(256);not null" json:"username"`
	Password string `gorm:"type:varchar(256);not null" json:"password"`
	Appid string `gorm:"type:varchar(256);not null"`
	CreateTime time.Time
	UpdateTime time.Time
}
