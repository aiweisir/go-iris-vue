package routes

import (
	"casbin-demo/models"
	"casbin-demo/services"
	"casbin-demo/supports"
	"casbin-demo/supports/jwts"
	"casbin-demo/utils"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func Registe(ctx iris.Context, u services.UserService) {
	user := models.User{}
	ctx.ReadJSON(&user)

	aes := utils.AESEncrypt([]byte(user.Password))
	user.Password = aes

	u.DoRegiste(&user)

	supports.Ok_(ctx, supports.Registe_success)
}

func Login(ctx iris.Context, u services.UserService) {
	user := new(models.User)
	ctx.ReadJSON(&user)

	mUser := u.DoLogin(user.Username)
	golog.Error(mUser)
	if mUser.Username == "" {
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
