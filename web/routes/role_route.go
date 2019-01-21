package routes

import (
	"go-iris/middleware/casbins"
	"go-iris/web/models"
	"go-iris/web/supports"
	"go-iris/web/supports/vo"
	"net/http"
	"strconv"
	"strings"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func RoleTable(ctx iris.Context) {
	//e := casbins.GetEnforcer()
	//golog.Info(e.GetAllRoles())

	pageNumber, err1 := ctx.URLParamInt("pageNumber")
	pageSize, err2 := ctx.URLParamInt("pageSize")
	sortName := ctx.URLParam("sortName")
	sortOrder := ctx.URLParam("sortOrder")
	golog.Infof("pageNumber=%d, pageSize=%d, sortName=%s, sortOrder=%s", pageNumber, pageSize, sortName, sortOrder)
	if err1 != nil || err2 != nil {
		ctx.Application().Logger().Errorf("查询角色列表参数解析错误. %s, %s", err1.Error(), err2.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	page := supports.Pagination{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		SortName:   sortName,
		SortOrder:  sortOrder,
	}
	page.PageSetting()

	rules, total, err := models.GetPaginationRoles(&page)
	if err != nil {
		ctx.Application().Logger().Errorf("查询角色列表错误. %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	ctx.JSON(vo.BootstrapTableVO{
		Total: total,
		Rows:  rules,
	})
}

// 创建角色
func CreateRole(ctx iris.Context) {
	//rule := new(supports.RoleDefine)
	rule := new(models.CasbinRule)
	if err := ctx.ReadJSON(&rule); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.ParseParamsFailur, nil)
	}

	e := casbins.GetEnforcer()
	ok := e.AddPolicy(rule.V0, rule.V1, rule.V2, rule.V3, rule.V4, rule.V5)
	if !ok {
		supports.Error(ctx, http.StatusInternalServerError, supports.RoleCreateFailur, nil)
	}
	supports.Ok_(ctx, supports.RoleCreateSuccess)
}

func UpdateRole(ctx iris.Context)  {
	role := new(models.CasbinRule)
	if err := ctx.ReadJSON(&role); err != nil {
		ctx.Application().Logger().Errorf("更新角色[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.OptionFailur, nil)
		return
	}

	effect, err := models.UpdateRoleById(role)
	if err != nil {
		ctx.Application().Logger().Errorf("更新角色[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, effect)
}

func DeleteRole(ctx iris.Context, rids string) {
	//groupDef := new(supports.GroupDefine)
	//if err := ctx.ReadJSON(groupDef); err != nil {
	//	supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, err.Error())
	//	return
	//}
	//
	//ok := true
	//e := casbins.GetEnforcer()
	//for _, v := range groupDef.Sub {
	//	if !e.DeleteRoleForUser(strconv.FormatInt(groupDef.Uid, 10), v) {
	//		ok = false
	//	}
	//}
	//
	//if !ok {
	//	supports.Error(ctx, http.StatusInternalServerError, supports.OptionFailur, nil)
	//	return
	//}
	//supports.Ok(ctx, supports.OptionSuccess, nil)

	ridList := strings.Split(rids, ",")
	if len(ridList) == 0 {
		ctx.Application().Logger().Error("删除角色错误, 参数不对.")
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	dRids := make([]int64, 0)
	for _, v := range ridList {
		if v == "" {
			continue
		}
		uid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ctx.Application().Logger().Error("删除角色错误, %s", err.Error())
			supports.Error(ctx, iris.StatusInternalServerError, supports.ParseParamsFailur, nil)
			return
		}
		dRids = append(dRids, uid)
	}

	effect, err := models.DeleteByRoles(dRids)
	if err != nil {
		ctx.Application().Logger().Error("删除角色错误, %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.DeleteRolesFailur, nil)
		return
	}
	supports.Ok(ctx, supports.DeleteRolesSuccess, effect)
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
