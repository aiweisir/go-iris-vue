package routes

import (
	"go-iris/web/models"
	"go-iris/web/services"
	"go-iris/web/supports"
	"net/http"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func AddOneProduct(ctx iris.Context, ds services.DemoService) {
	demo := new(models.Demo)
	if err := ctx.ReadJSON(demo); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, err.Error())
		return
	}

	golog.Info(demo)
	demo.CreateDate = time.Now()
	err := ds.AddOneProduct(demo)
	if err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, err.Error())
		return
	}
	supports.Ok_(ctx, supports.Option_success)
}

func GetOneProduct(ctx iris.Context, ds services.DemoService, pid int64) {
	demo := new(models.Demo)
	demo.Pid = pid
	_, err := ds.GetOneProduct(demo)
	if err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, nil)
		return
	}
	supports.Ok(ctx, supports.Option_success, demo)
}
