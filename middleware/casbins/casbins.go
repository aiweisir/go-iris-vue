package casbins

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbins".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local") // Your driver and data source.
	a = xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/") // Your driver and data source.
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	e = casbin.NewEnforcer("conf/rbac_model.conf", a)
)

// Casbin is the casbins services which contains the casbins enforcer.
//type Casbin struct {
//	Enforcer *casbins.Enforcer
//}

// New returns the casbins service which receives a casbins enforcer.
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

// ServeHTTP is the iris compatible casbins handler which should be passed to specific routes or parties.
// Usage:
// [...]
// app.Get("/dataset1/resource1", casbinMiddleware.ServeHTTP, myHandler)
// [...]
//func CheckPermissions(ctx context.Context) bool {
//
//	// jwt token拦截
//	//jwts.ConfigJWT().Serve(ctx)
//	//a := ctx.Values().Get(jwts.DefaultContextKey).(*jwt.Token)
//	//golog.Infof("req set values jwt key =%s", a.Claims)
//
//	 //casbin权限拦截
//	yes := e.Enforce("alice", "", "get", ".*")
//	golog.Infof("path= %s, casbincheck= %t\n", "", yes)
//	if !yes {
//		supports.Unauthorized(ctx, supports.Permissions_less, nil)
//		ctx.StopExecution()
//		return false
//	}
//
//	return true
//	//ctx.Next()
//}

// Wrapper is the router wrapper, prefer this method if you want to use casbins to your entire iris application.
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
