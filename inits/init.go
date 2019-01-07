package inits

import (
	"go-iris/middleware/casbins"
	"go-iris/utils"
	"go-iris/web/db"
	"go-iris/web/db/mappers"
	"go-iris/web/models"
	"go-iris/web/routes/dispatch"
	"go-iris/web/services"
	"strconv"
	"time"

	"github.com/kataras/golog"
)

func init() {
	initRootUser()
	initServices()
}

const (
	username = "root"
	password = "123456"
)
// 暴露给casbin
var RootID string

func initRootUser() {
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
	RootID = strconv.FormatInt(newRoot.Id, 10)

	// add policy for root
	p := casbins.GetEnforcer().AddPolicy(utils.FmtRolePrefix(newRoot.Id), "/*", "ANY", ".*")
	if !p {
		golog.Fatalf("初始化用户[%s]权限失败。%s", username, err.Error())
	}

	golog.Infof("@@@ Create Root User and add-permissions OK, Effect %d row.", effRow)
}

func initServices() {
	golog.Info("@@@ Inject all service")

	dispatch.Register(
		services.NewUserService(mappers.NewUserMapper()),
		services.NewDemoService(mappers.NewDemoMapper()),
	)
}
