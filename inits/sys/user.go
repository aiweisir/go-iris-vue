package sys

import (
	"go-iris/middleware/casbins"
	"go-iris/utils"
	"go-iris/web/db"
	"go-iris/web/models"
	"strconv"
	"time"

	"github.com/kataras/golog"
)

const (
	username = "root"
	password = "123456"
)

// 检查超级用户是否存在
func CheckRootExit() bool {
	e := db.MasterEngine()
	// root is existed?
	exit, err := e.Exist(&models.User{Username: username})
	if err != nil {
		golog.Fatalf("@@@ When check Root User is exited? happened error. %s", err.Error())
	}
	if exit {
		golog.Info("@@@ Root User is existed.")

		// 初始化rbac_model
		r := models.User{Username: username}
		if exit, _ := e.Get(&r); exit {
			casbins.SetRbacModel(strconv.FormatInt(r.Id, 10))
			CreateSystemRole()
		}
	}
	return exit
}

func CreateRoot() {
	newRoot := models.User{
		Username:   username,
		Password:   utils.AESEncrypt([]byte(password)),
		CreateTime: time.Now(),
	}

	e := db.MasterEngine()
	if _, err := e.Insert(&newRoot); err != nil {
		golog.Fatalf("@@@ When create Root User happened error. %s", err.Error())
	}
	rooId := strconv.FormatInt(newRoot.Id, 10)
	casbins.SetRbacModel(rooId)

	addAllpolicy(rooId)
}

func addAllpolicy(rooId string) {
	// add policy for root
	//p := casbins.GetEnforcer().AddPolicy(utils.FmtRolePrefix(newRoot.Id), "/*", "ANY", ".*")
	e := casbins.GetEnforcer()
	p := e.AddPolicy(rooId, "/*", "ANY", ".*", "", "", "", "", "", "超级用户")
	if !p {
		golog.Fatalf("初始化用户[%s]权限失败.", username)
	}

	//
	for _, v := range Components {
		e.AddGroupingPolicy(rooId, v[0])
	}
}
