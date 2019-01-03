package routes

import (
	"casbin-demo/middleware/jwts"
	"casbin-demo/models"
	"casbin-demo/services"
	"casbin-demo/supports"
	"casbin-demo/utils"

	"github.com/kataras/iris"
)

func Registe(ctx iris.Context, u services.UserService) {
	user := new(models.User)
	ctx.ReadJSON(&user)

	err := u.DoRegiste(user)
	if err != nil {
		ctx.Application().Logger().Error("用户[%s]注册失败。%s", user.Username, err)
		supports.Error(ctx, iris.StatusInternalServerError, supports.Registe_failur, nil)
	}else {
		supports.Ok_(ctx, supports.Registe_success)
	}
}

func Login(ctx iris.Context, u services.UserService) {
	user := new(models.User)
	ctx.ReadJSON(&user)

	mUser := new(models.User)
	mUser.Username = user.Username
	has, err := u.DoLogin(mUser)
	//golog.Error(mUser)
	if err != nil {
		ctx.Application().Logger().Error("用户[%s]登录失败。%s", user.Username, err)
		supports.Error(ctx, iris.StatusInternalServerError, supports.Login_failur, nil)
		return
	}

	if !has { // 用户名不正确
		supports.Unauthorized(ctx, supports.Username_failur, nil)
		return
	}

	ckPassword := utils.CheckPWD(user.Password, mUser.Password)
	if !ckPassword {
		supports.Unauthorized(ctx, supports.Password_failur, nil)
		return
	}

	token, err := jwts.GenerateToken(user.Username, user.Password);
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录，生成token出错。%s", user.Username, err)
		supports.Error(ctx, iris.StatusInternalServerError, supports.Token_create_failur, nil)
		return
	}

	supports.Ok(ctx, supports.Login_success, token)
	return
}
