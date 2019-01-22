package main

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	_ "go-iris/inits"
	"go-iris/inits/parse"
	"go-iris/middleware/preset"
	"go-iris/route_Controller"
)
// $ go get github.com/casbins/casbins
// $ go run main.go
func main() {
	app := iris.New()
	preset.PreSettring(app)
	route_Controller.Hub(app)

	app.RegisterView(iris.HTML("resources", ".html"))
	app.StaticWeb("/static", "resources/static") // 设置静态资源
	golog.Info()
	app.Run(iris.Addr(":8088"), iris.WithConfiguration(parse.C))
}
/*
func newApp() *iris.Application {
	app := iris.New()
	preset.PreSettring(app)
	route_Controller.Hub(app)
	return app
}
*/
