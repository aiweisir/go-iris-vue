package main

import (
	_ "go-iris/inits"
	"go-iris/inits/parse"
	"go-iris/middleware"
	"go-iris/web/routes"
	"go-iris/web/routes/dispatch"
	"go-iris/web/supports"

	"github.com/iris-contrib/middleware/cors"
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

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		Debug:            true,
		//AllowCredentials: true,
	})

	/* 定义路由 */
	main := app.Party("/", crs).AllowMethods(iris.MethodOptions)
	main.Use(middleware.ServeHTTP)
	//main := app.Party("/")
	home := main.Party("/")
	{
		home.Get("/", func(ctx iris.Context) { // 首页模块
			ctx.View("index.html")
		})
		home.Get("/sysMenu", dispatch.Handler(routes.DynamicMenu)) // 获取动态菜单
	}

	// 用户API模块
	user := main.Party("/user")
	{
		//p.Use(middleware.BasicAuth)
		user.Post("/registe", dispatch.Handler(routes.Registe))
		user.Post("/login", dispatch.Handler(routes.Login))
	}

	// 权限API模块
	admin := main.Party("/admin")
	{
		users := admin.Party("/users")
		{
			users.Get("/", dispatch.Handler(routes.UserTable)) // 用户列表
		}

		role := admin.Party("/role")
		{
			role.Post("/", dispatch.Handler(routes.CreateRole))   // 添加角色
			role.Get("/", dispatch.Handler(routes.AllRoleOfUser)) // 获取所有角色
			role.Delete("/", dispatch.Handler(routes.DeleteRole)) // 删除角色
		}

		permissions := admin.Party("/permissions")
		{
			permissions.Post("/permissions", dispatch.Handler(routes.RelationUserRole)) // 给角色添加权限
		}
	}

	// demo测试API模块
	demo := main.Party("/demo")
	{
		demo.Get("/{pid:long}", dispatch.Handler(routes.GetOneProduct))
		demo.Put("/", dispatch.Handler(routes.AddOneProduct))
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
		//middleware.ServeHTTP
	)

	// ---------------------- 定义错误处理 ------------------------
	app.OnErrorCode(iris.StatusNotFound, customLogger, func(ctx iris.Context) {
		supports.Error(ctx, iris.StatusNotFound, supports.NotFound, nil)
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
