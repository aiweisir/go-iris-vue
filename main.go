package main

import (
	"casbin-demo/conf/parse"
	"casbin-demo/middleware"
	"casbin-demo/routes"
	"casbin-demo/routes/dispatch"
	"casbin-demo/supports"

	"github.com/kataras/iris"

	//cm "github.com/iris-contrib/middleware/casbins"

	// Inject all service
	_ "casbin-demo/services"
	// Init all configuration
	_ "casbin-demo/conf/parse"

	"github.com/kataras/iris/middleware/logger"
	rcover "github.com/kataras/iris/middleware/recover"
)

// $ go get github.com/casbins/casbins
// $ go run main.go

func main() {
	app := newApp()
	app.Logger().SetLevel(parse.O.LogLevel)
	app.RegisterView(iris.HTML("resources", ".html").Reload(true))
	app.StaticWeb("/static", "resources/static") // 设置静态资源

	app.Run(
		iris.Addr(":8080"),
		iris.WithConfiguration(parse.C))
}

func newApp() *iris.Application {
	app := iris.New()
	registerMiddlewareAndDefError(app)

	app.PartyFunc("/", func(p iris.Party) {
		p.Get("/", func(ctx iris.Context) {
			ctx.View("index.html")
		})

		user := p.Party("/user")
		{
			//	// Add the basic authentication(admin:password) middleware
			//	// for the /movies based requests.
			//p.Use(middleware.BasicAuth)
			user.Post("/registe", dispatch.Handler(routes.Registe))
			user.Post("/login", dispatch.Handler(routes.Login))
		}
		user.Post("/s", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"a": 1,
				"b": 2,
			})
		})
	})
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
