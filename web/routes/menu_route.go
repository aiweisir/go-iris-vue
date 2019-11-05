package routes

import (
	"go-iris/web/models"
	"go-iris/web/supports"
	"go-iris/web/supports/vo"

	"github.com/kataras/iris/v12"
)

func MenuTable(ctx iris.Context) {
	page, err := supports.NewPagination(ctx)
	if err != nil {
		ctx.Application().Logger().Errorf("查询菜单列表参数解析错误. %s", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	menus, total, err := models.GetPaginationMenus(page)
	if err != nil {
		ctx.Application().Logger().Errorf("查询菜单列表错误. %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	ctx.JSON(vo.BootstrapTableVO{
		Total: total,
		Rows:  menus,
	})
}

// 修改角色权限

// 或给用户设置权限
