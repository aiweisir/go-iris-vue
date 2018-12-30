package supports

import "github.com/kataras/iris"

const (
	// key定义
	CODE string = "code"
	MSG  string = "msg"
	DATA string = "data"

	// msg define
	Success                = "成功"
	Option_success  string = "操作成功"
	Registe_success  string = "操作成功"
	Login_success   string = "登录成功"
	Username_failur string = "用户名错误"
	Password_failur string = "密码错误"
	Token_failur    string = "token错误"
	Token_create_failur    string = "生成token错误"
	Not_found       string = "您请求的url不存在"

	// value define

)

// 200 define
func Ok_(ctx iris.Context, msg string) {
	Ok(ctx, msg, nil)
}

func Ok(ctx iris.Context, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(&iris.Map{
		CODE: iris.StatusOK,
		MSG:  msg,
		DATA: data,
	})
}

// 401 error define
func Unauthorized(ctx iris.Context, msg string, data interface{}) {
	unauthorized := iris.StatusUnauthorized

	ctx.StatusCode(unauthorized)
	ctx.JSON(&iris.Map{
		CODE: unauthorized,
		MSG:  msg,
		DATA: data,
	})
}

// common error define
func Error(ctx iris.Context, status int, msg string, data interface{}) {
	ctx.StatusCode(status)
	ctx.JSON(&iris.Map{
		CODE: status,
		MSG:  msg,
		DATA: data,
	})
}
