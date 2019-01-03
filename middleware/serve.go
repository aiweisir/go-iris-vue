package middleware

import (
	"casbin-demo/conf/parse"
	"casbin-demo/middleware/casbins"
	"casbin-demo/middleware/jwts"
	"casbin-demo/supports"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"github.com/kataras/iris/context"
)

type Middleware struct {

}

func Serve(ctx context.Context) {
	return
	path := ctx.Path()
	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) || strings.Contains(path, "/static") {
		ctx.Next()
		return
	}
	//
	// jwt token拦截
	jwts.ConfigJWT().Serve(ctx)
	token := ctx.Values().Get(jwts.DefaultContextKey)
	if token == nil {
		//supports.Unauthorized(ctx, supports.Token_failur, nil)
		//ctx.StopExecution()
		return
	}

	golog.Infof("jwt key =%v", token)
	golog.Infof("req set values of jwt key =%s", token.(*jwt.Token).Claims)

	// casbin权限拦截
	yes := casbins.GetEnforcer().Enforce("alice", path, "get", ".*")
	golog.Infof("path= %s, casbincheck= %t\n", path, yes)
	//yes := true

	//yes2 := casbins.CheckPermissions(ctx)

	if !yes {
		supports.Unauthorized(ctx, supports.Permissions_less, nil)
		ctx.StopExecution()
		return
	}

	ctx.Next()
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
		if reqPath == v {
			return true
		}
	}
	return false
}
