package casbins

import (
	"casbin-demo/conf/parse"
	"net/http"
	"strings"

	"github.com/casbin/gorm-adapter"
	"github.com/kataras/iris/context"

	"github.com/casbin/casbin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local") // Your driver and data source.
	a = gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/") // Your driver and data source.
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	e = casbin.NewEnforcer("conf/rbac_model.conf", a)
)

// Casbin is the casbins services which contains the casbin enforcer.
//type Casbin struct {
//	Enforcer *casbin.Enforcer
//}

// New returns the casbins service which receives a casbin enforcer.
//
// Adapt with its `Wrapper` for the entire application
// or its `ServeHTTP` for specific routes or parties.
//func New() *Casbin {
//	return &Casbin{Enforcer: e}
//}

// 获取Enforcer
func GetEnforcer() *casbin.Enforcer {
	//e.AddPermissionForUser()
	return e
}

// ServeHTTP is the iris compatible casbin handler which should be passed to specific routes or parties.
// Usage:
// [...]
// app.Get("/dataset1/resource1", casbinMiddleware.ServeHTTP, myHandler)
// [...]
func Serve(ctx context.Context) {
	path := ctx.Path()
	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) || strings.Contains(path, "/static") {
		ctx.Next()
		return
	}

	// casbin权限拦截
	//c := e.Enforce("alice", path, "get", ".*")
	//fmt.Printf("path= %s, casbincheck= %t\n", path, c)
	//if !c {
	//	ctx.StatusCode(http.StatusForbidden) // Status Forbiden
	//	ctx.JSON(utils.Error(iris.StatusForbidden, "权限不足", nil))
	//	//ctx.StopExecution()
	//	return
	//}

	// jwt token拦截
	//jwts.ConfigJWT().Serve(ctx)
	ctx.Next()
}

func filter() {

}

/**
return
	true:则跳过不需验证，如登录接口等...
	false:需要进一步验证
 */
func checkURL(reqPath string) bool {
	//config := iris.YAML("conf/app.yml")
	//ignoreURLs := config.GetOther()["ignoreURLs"].([]interface{})
	for _, v := range parse.O.IgnoreURLs {
		if (reqPath == v) {
			return true
		}
	}
	return false
}

// Wrapper is the router wrapper, prefer this method if you want to use casbin to your entire iris application.
// Usage:
// [...]
// app.WrapRouter(casbinMiddleware.Wrapper())
// app.Get("/dataset1/resource1", myHandler)
// [...]
func Wrapper() func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		//if !c.Check(r) {
		//	w.WriteHeader(http.StatusForbidden)
		//	w.Write([]byte("403 Forbidden"))
		//	return
		//}
		router(w, r)
	}
}
