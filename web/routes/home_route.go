package routes

import (
	"go-iris/middleware/jwts"
	"go-iris/web/models"
	"go-iris/web/supports"

	"github.com/kataras/iris/v12"
)

func DynamicMenu(ctx iris.Context) {
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		return
	}

	menuTree := models.DynamicMenuTree(user.Id)
	supports.Ok(ctx, supports.OptionSuccess, menuTree)
}
