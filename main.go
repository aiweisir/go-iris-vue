package main

import (
	_ "go-iris/inits"
	"go-iris/inits/parse"
	"go-iris/middleware"
	"go-iris/web/routes"
	"go-iris/web/routes/dispatch"
	"go-iris/web/supports"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	rcover "github.com/kataras/iris/middleware/recover"
)

// $ go get github.com/casbins/casbins
// $ go run main.go
func main() {
	app := newApp()
	app.Logger().SetLevel(parse.O.LogLevel)
	app.RegisterView(iris.HTML("resources", ".html"))
	app.StaticWeb("/static", "resources/static") // 设置静态资源

	golog.Info()
	app.Run(iris.Addr(":8088"), iris.WithConfiguration(parse.C))
}

func newApp() *iris.Application {
	app := iris.New()
	registerMiddlewareAndDefError(app)

	app.PartyFunc("/", func(home iris.Party) {
		home.Get("/", func(ctx iris.Context) {
			ctx.View("index.html")
		})
		user := home.Party("/user")
		{
			//p.Use(middleware.BasicAuth)
			user.Post("/registe", dispatch.Handler(routes.Registe))
			user.Post("/login", dispatch.Handler(routes.Login))

			// permission manage api
			manage := user.Party("/manage")
			{
				manage.Post("/role", dispatch.Handler(routes.AddRole))
				//manage.Delete("/role", dispatch.Handler(nil))
				manage.Post("/permissions", dispatch.Handler(routes.AddPermissions))
				//manage.Delete("/permissions", dispatch.Handler(nil))
			}
		}
		user.Post("/s", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"a": 1,
				"b": 2,
			})
		})
	})

	demo := app.Party("/demo")
	{
		demo.Get("/{pid:long}", dispatch.Handler(routes.GetOneProduct))
		demo.Put("/", dispatch.Handler(routes.AddOneProduct))
	}

	a := app.Party("/a")
	{
		a.Get("/a1", func(ctx iris.Context) {
			ctx.JSON("/a/a1, get")
		})
		a.Post("/a2", func(ctx iris.Context) {
			ctx.JSON("/a/a2, post")
		})
	}

	return app
}

// 注册中间件、定义错误处理
func registerMiddlewareAndDefError(app *iris.Application) {
	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(
		rcover.New(),
		customLogger,
		middleware.ServeHTTP)

	// ---------------------- 定义错误处理 ------------------------
	app.OnErrorCode(iris.StatusNotFound, customLogger, func(ctx iris.Context) {
		supports.Error(ctx, iris.StatusNotFound, supports.Not_found, nil)
	})
	//app.OnErrorCode(iris.StatusForbidden, customLogger, func(ctx iris.Context) {
	//	ctx.JSON(utils.Error(iris.StatusForbidden, "权限不足", nil))
	//})
	//捕获所有http错误:
	//app.OnAnyErrorCode(customLogger, func(ctx iris.Context) {
	//	//这应该被添加到日志中，因为`logger.Config＃MessageContextKey`
	//	ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	//	ctx.JSON(utils.Error(500, "服务器内部错误", nil))
	//})
}

// -------------------------------------------------------------------------
// -------------------------------------------------------------------------
// -------------------------------------------------------------------------
