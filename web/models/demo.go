package models

import "time"

type Demo struct {
	Pid         int64     `xorm:"pk autoincr INT(10) notnull" json:"pid"`
	ProductCode string    `xorm:"notnull" json:"productCode"`
	ProductName string    `xorm:"notnull" json:"productName"`
	Number      int       `xorm:"notnull" json:"number"`
	CreateDate  time.Time `xorm:"notnull" json:"createDate"`
}
