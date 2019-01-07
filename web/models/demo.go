package models

type Demo struct {
	Pid         int64    `xorm:"pk autoincr INT(10) notnull" json:"pid"`
	ProductCode string `xorm:"notnull" json:"productCode"`
	ProductName string `xorm:"notnull" json:"productName"`
	Number      string `xorm:"notnull" json:"number"`
	CreateDate  string `xorm:"notnull" json:"createDate"`
}
