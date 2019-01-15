package inits

import (
	"go-iris/inits/parse"
	"go-iris/inits/sys"
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

	sys.CreateSystemRole()

}


