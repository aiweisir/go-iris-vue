package main

import (
	"casbin-demo/casbins"
	"casbin-demo/conf/parse"
	"casbin-demo/routes"
	"casbin-demo/routes/dispatch"
	"casbin-demo/supports"

	"github.com/kataras/iris"

	//cm "github.com/iris-contrib/middleware/casbin"

	// 注入路由
	_ "casbin-demo/services"
	// 初始化数据库配置
	_ "casbin-demo/conf/parse"

	"github.com/kataras/iris/middleware/logger"
	rcover "github.com/kataras/iris/middleware/recover"
)

// $ go get github.com/casbin/casbin
// $ go run main.go

func main() {
	c := iris.YAML("conf/app.yml")
	parse.InitOtherConfig(&c)

	app := newApp()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("resources", ".html").Reload(true))
	app.StaticWeb("/static", "resources/static") // 设置静态资源

	app.Run(
		iris.Addr(":8080"),
		iris.WithConfiguration(c))
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
	})
	return app
}

// 注册中间件、定义错误处理
func registerMiddlewareAndDefError(app *iris.Application) {
	// ---------------------- 注册中间件 ------------------------
	app.Use(rcover.New())

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
	app.Use(customLogger, casbins.Serve)

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
