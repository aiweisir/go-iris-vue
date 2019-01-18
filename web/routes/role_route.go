package routes

import (
	"go-iris/middleware/casbins"
	"go-iris/web/supports"
	"net/http"
	"strconv"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func RoleTable(ctx iris.Context) {
	e := casbins.GetEnforcer()
	golog.Info(e.GetAllRoles())

	//supports.Ok(ctx, supports.OptionSuccess, roles)
}

// 创建角色
func CreateRole(ctx iris.Context) {
	rule := new(supports.RoleDefine)
	if err := ctx.ReadJSON(rule); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.Role_create_failur, nil)
	}

	e := casbins.GetEnforcer()
	ok := e.AddPolicy(rule.Sub, rule.Obj, rule.Act, rule.Suf, rule.RoleName)
	if !ok {
		supports.Error(ctx, http.StatusInternalServerError, supports.Role_create_failur, nil)
	}
	supports.Ok_(ctx, supports.OptionSuccess)
}

func DeleteRole(ctx iris.Context) {
	groupDef := new(supports.GroupDefine)
	if err := ctx.ReadJSON(groupDef); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	ok := true
	e := casbins.GetEnforcer()
	for _, v := range groupDef.Sub {
		if !e.DeleteRoleForUser(strconv.FormatInt(groupDef.Uid, 10), v) {
			ok = false
		}
	}

	if !ok {
		supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, nil)
}

// 给用户指定角色
func RelationUserRole(ctx iris.Context) {
	groupDef := new(supports.GroupDefine)
	if err := ctx.ReadJSON(groupDef); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	// TODO 校验前端的角色是否正确，和数据库的所有角色比较

	ok := true
	e := casbins.GetEnforcer()
	for _, v := range groupDef.Sub {
		// 给目标用户添加角色
		if !e.AddGroupingPolicy(strconv.FormatInt(groupDef.Uid, 10), v) {
			ok = false
		}
	}

	if !ok {
		supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok_(ctx, supports.OptionSuccess)
}
