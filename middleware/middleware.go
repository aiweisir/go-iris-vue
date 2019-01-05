package middleware

import (
	"casbin-demo/inits/parse"
	"casbin-demo/middleware/casbins"
	"casbin-demo/middleware/jwts"
	"strings"

	"github.com/kataras/iris/context"
)

type Middleware struct {
}

func ServeHTTP(ctx context.Context) {
	path := ctx.Path()
	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) || strings.Contains(path, "/static") {
		ctx.Next()
		return
	}

	// jwt token拦截
	token := jwts.Serve(ctx)
	if token == nil {
		//supports.Unauthorized(ctx, supports.Token_failur, nil)
		//ctx.StopExecution()
		return
	}

	// casbin权限拦截
	ok := casbins.CheckPermissions(ctx, token)
	if !ok {
		//supports.Unauthorized(ctx, supports.Permissions_less, nil)
		//ctx.StopExecution()
		return
	}

	// Pass to real API
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
