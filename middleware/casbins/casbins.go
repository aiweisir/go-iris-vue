package casbins

import (
	"fmt"
	"go-iris/inits/parse"
	"go-iris/web/db"
	"go-iris/web/supports"
	"net/http"
	"strconv"
	"sync"

	"github.com/casbin/casbin"

	//"github.com/casbin/xorm-adapter"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var (
	adt *Adapter // Your driver and data source.
	e   *casbin.Enforcer

	adtLook sync.Mutex
	eLook sync.Mutex

	rbacModel string
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

func SetRbacModel(rootID string) {
	rbacModel = fmt.Sprintf(`
[request_definition]
r = sub, obj, act, suf

[policy_definition]
p = sub, obj, act, suf

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.suf, p.suf) && regexMatch(r.act, p.act) || r.sub == "%s"
`, rootID)
}

// 获取Enforcer
func GetEnforcer() *casbin.Enforcer {
	if e != nil {
		e.LoadPolicy()
		return e
	}
	eLook.Lock()
	defer eLook.Unlock()
	if e != nil {
		e.LoadPolicy()
		return e
	}

	m := casbin.NewModel(rbacModel)
	//m.AddDef("r", "r", "sub, obj, act, suf")
	//m.AddDef("p", "p", "sub, obj, act, suf")
	//m.AddDef("g", "g", "_, _")
	//m.AddDef("e", "e", "some(where (p.eft == allow))")
	//m.AddDef("m", "m", `g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.suf, p.suf) && regexMatch(r.act, p.act) || r.sub == "1"`)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	// TODO use go-bindata fill
	//e = casbin.NewEnforcer("conf/rbac_model.conf", singletonAdapter())
	e = casbin.NewEnforcer(m, singleAdapter())
	e.EnableLog(true)
	return e
}

func singleAdapter() *Adapter {
	if adt != nil {
		return adt
	}
	adtLook.Lock()
	defer adtLook.Unlock()
	if adt != nil {
		return adt
	}

	master := parse.DBConfig.Master
	url := db.GetConnURL(&master)
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbins".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local") // Your driver and data source.
	adt = NewAdapter(master.Dialect, url, true) // Your driver and data source.
	return adt
}

// ServeHTTP is the iris compatible casbins handler which should be passed to specific routes or parties.
// Usage:
// [...]
// app.Get("/dataset1/resource1", casbinMiddleware.ServeHTTP, myHandler)
// [...]
func CheckPermissions(ctx context.Context, token *jwt.Token) bool {
	mapClaims := (token.Claims).(jwt.MapClaims)
	id, ok := mapClaims["id"].(float64)
	golog.Infof("*** MapClaims=%v, id=%f, isOK=%t", mapClaims, id, ok)
	if !ok {
		supports.Error(ctx, iris.StatusInternalServerError, supports.Token_parse_failur, nil)
		return false
	}

	uid := strconv.Itoa(int(id))
	yes := GetEnforcer().Enforce(uid, ctx.Path(), ctx.Method(), ".*")
	//golog.Infof("*** uid=%d, Path=%s, Method=%s, Permission=%t", int(id), ctx.Path(), ctx.Method(), yes)
	if !yes {
		supports.Unauthorized(ctx, supports.Permissions_less, nil)
		ctx.StopExecution()
		return false
	}

	ctx.Values().Set("uid", uid)
	return true
	//ctx.Next()
}

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
