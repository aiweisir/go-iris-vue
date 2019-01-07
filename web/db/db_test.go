package db

import (
	"go-iris/inits/parse"
	"testing"
)

func Test_p(t *testing.T) {
	parse.ParseDBSetting()
}

func Test_r(t *testing.T) {
	master := MasterEngine()
	//u := models.User{
	//	Username:   "yhm7",
	//	Password:   "12345",
	//	CreateTime: time.Now(),
	//	UpdateTime: time.Now(),
	//}


	master.Ping()

	//插入记录
	//t.Log(master.NewRecord(u))
	//master.Create(&u)
	//t.Log(master.NewRecord(u))
	////root := models.User{
	////	Username: "root22",
	////	Password: "root2",
	////	CreateTime: time.Now(),
	////	UpdateTime: time.Now(),
	////}
	////master.Create(root)
	////master.Commit()
	//
	//u2 := []models.User{}
	//// 自定义sql查询
	//master.Raw("select * from user").Scan(&u2)
	//t.Log(u2)
}
