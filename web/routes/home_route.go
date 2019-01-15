package routes

import (
	"go-iris/middleware/casbins"
	"go-iris/web/supports"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func DynamicMenu(ctx iris.Context)  {
	uid := ctx.Values().Get("uid").(string)

	golog.Info("===>", uid)
	res := casbins.GetAllResourcesByUID(uid)
	supports.Ok(ctx, supports.Option_success, res)
}




