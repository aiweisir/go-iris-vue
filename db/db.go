package db

import (
	"casbin-demo/conf/parse"
	"fmt"
	"sync"

	"github.com/kataras/golog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	masterDB *gorm.DB
	slaveDB  *gorm.DB
	lock     sync.Mutex
)

// 主库，单例
func MasterDB() *gorm.DB {
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return "prefix_" + defaultTableName;
	//}

	if masterDB != nil {
		return masterDB
	}

	lock.Lock()
	defer lock.Unlock()

	if masterDB != nil {
		return masterDB
	}

	master := parse.DBConfig.Master
	db, err := gorm.Open(master.Dialect, getConnURL(&master))
	if err != nil {
		golog.Fatalf("@@@ Instance Master DB error!! %s", err)
		return nil
	}
	// 禁用 表名 启用复数
	db.SingularTable(true)
	masterDB = db

	return db
}

// 从库，单例
func SlaveDB() *gorm.DB {
	if slaveDB != nil {
		return slaveDB
	}

	lock.Lock()
	defer lock.Unlock()

	if slaveDB != nil {
		return slaveDB
	}

	slave := parse.DBConfig.Slave
	db, err := gorm.Open(slave.Dialect, getConnURL(&slave))
	if err != nil {
		golog.Fatalf("@@@ Instance Slave DB error!! %s", err)
		return nil
	}
	// 禁用 表名 启用复数
	db.SingularTable(true)
	slaveDB = db

	return db
}

// 获取数据库连接的url
// true：master主库
func getConnURL(info *parse.DBConfigInfo) (url string) {
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	url = fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%t",
		info.User,
		info.Password,
		info.Database,
		info.Charset,
		info.ParseTime)
	golog.Info(url)
	return
}
