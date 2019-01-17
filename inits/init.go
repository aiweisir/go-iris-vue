package inits

import (
	"go-iris/inits/parse"
	"go-iris/inits/sys"
	"go-iris/web/models"

	"github.com/kataras/golog"
)

func init() {
	parse.AppOtherParse()
	parse.DBSettingParse()

	initRootUser()
}

func initRootUser() {
	// root is existed?
	if sys.CheckRootExit() {
		return
	}

	// create root user
	sys.CreateRoot()

	ok := sys.CreateSystemRole()
	if ok {
		addRoleMenu()
	}

}

func addRoleMenu()  {
	// 添加role-menu关系
	rMenus := []*models.RoleMenu {
		{Rid:68, Mid:2},
		{Rid:68, Mid:3},
		{Rid:68, Mid:4},
		{Rid:68, Mid:5},
	}
	effect, err := models.CreateRelationRoleMenu(rMenus...)
	if err != nil {
		golog.Fatalf("**@@@@@@@@@@@0> %d, %s", effect, err.Error())
	}
	golog.Infof("@@@@@@@@@-> %s, %s", effect, err.Error())
}


