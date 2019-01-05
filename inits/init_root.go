package inits

import (
	"casbin-demo/db"
	"casbin-demo/middleware/casbins"
	"casbin-demo/models"
	"casbin-demo/utils"
	"strconv"
	"time"

	"github.com/kataras/golog"
)

const (
	username = "root"
	password = "123456"
)

// Initial Root User
func init() {
	e := db.MasterEngine()

	// root is existed?
	exit, err := e.Exist(&models.User{Username: username})
	if err != nil {
		golog.Fatalf("@@@ When check Root User is exited? happened error. %s", err.Error())
	}
	if exit {
		golog.Info("@@@ Root User is existed.")
		return
	}

	// create root user
	newRoot := models.User{
		Username:   username,
		Password:   utils.AESEncrypt([]byte(password)),
		CreateTime: time.Now(),
	}
	effRow, err := e.Insert(&newRoot)
	if err != nil {
		golog.Fatalf("@@@ When create Root User happened error. %s", err.Error())
	}

	// add policy for root
	mRoot := new(models.User)
	mRoot.Username = username
	has, err := e.Get(mRoot)
	if err != nil {
		golog.Fatalf("初始化用户[%s]权限时，查询数据库失败。%s", username, err.Error())
	}
	if has {
		p := casbins.GetEnforcer().AddPolicy(strconv.FormatInt(mRoot.Id, 10), "/*", "ANY", ".*")
		if !p {
			golog.Fatalf("初始化用户[%s]权限失败。%s", username, err.Error())
		}
	}

	golog.Infof("@@@ Create Root User and add-permissions OK, Effect %d row.", effRow)
}
