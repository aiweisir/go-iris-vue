package inits

import (
	"go-iris/middleware/casbins"
	"go-iris/utils"
	"go-iris/web/db"
	"go-iris/web/models"
	"strconv"
	"time"

	"github.com/kataras/golog"
)

func init() {
	initRootUser()
	//initServices()
}

const (
	username = "root"
	password = "123456"
)

func initRootUser() {
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
		}
		return
	}

	// create root user
	newRoot := models.User{
		Username:   username,
		Password:   utils.AESEncrypt([]byte(password)),
		CreateTime: time.Now(),
	}
	effRow, err := e.Insert(&newRoot)
	golog.Errorf("----> %s, %v, %d", newRoot.Id, newRoot, effRow)
	if err != nil {
		golog.Fatalf("@@@ When create Root User happened error. %s", err.Error())
	}
	casbins.SetRbacModel(strconv.FormatInt(newRoot.Id, 10))

	// add policy for root
	p := casbins.GetEnforcer().AddPolicy(utils.FmtRolePrefix(newRoot.Id), "/*", "ANY", ".*")
	if !p {
		golog.Fatalf("初始化用户[%s]权限失败。%s", username, err.Error())
	}

	golog.Infof("@@@ Create Root User and add-permissions OK, Effect %d row.", effRow)
}

func initServices() {
	golog.Info("@@@ Inject all service")

	//dispatch.Register(
	//	services.NewUserService(mappers.NewUserMapper()),
	//	services.NewDemoService(mappers.NewDemoMapper()),
	//)
}
