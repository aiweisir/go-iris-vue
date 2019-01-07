package routes

import (
	"go-iris/web/models"
	"go-iris/web/services"
	"go-iris/web/supports"
	"net/http"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func AddOneProduct(ctx iris.Context, ds services.DemoService) {
	demo := new(models.Demo)
	ctx.ReadJSON(demo)

	golog.Info(demo)
	err := ds.AddOneProduct(demo)
	if err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, nil)
	}
	supports.Ok_(ctx, supports.Option_success)
}

func GetOneProduct(ctx iris.Context, ds services.DemoService, pid int64) {
	pid2, e := ctx.Params().GetInt("pid")
	golog.Infof("query pid2 err. %d, %s", pid2, e.Error())
	golog.Infof("pid=%d", pid)

	demo := new(models.Demo)
	demo.Pid = pid
	_, err := ds.GetOneProduct(demo)
	if err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, nil)
	}
	supports.Ok(ctx, supports.Option_success, demo)
}
