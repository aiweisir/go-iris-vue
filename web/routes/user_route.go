package routes

import (
	"go-iris/middleware/jwts"
	"go-iris/utils"
	"go-iris/web/models"
	"go-iris/web/supports"
	"go-iris/web/supports/vo"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

func UserHub(app *iris.Application) {

}

func Registe(ctx iris.Context) {
	user := new(models.User)
	ctx.ReadJSON(&user)

	user.Password = utils.AESEncrypt([]byte(user.Password))
	user.Enable = 1
	user.CreateTime = time.Now()

	effect, err := models.CreateUser(user)
	if effect <= 0 || err != nil {
		ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Username, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.RegisteFailur, nil)
	} else {
		supports.Ok_(ctx, supports.RegisteSuccess)
	}
}

func Login(ctx iris.Context) {
	user := new(models.User)
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.LoginFailur, nil)
		return
	}

	mUser := new(models.User)
	mUser.Username = user.Username
	has, err := models.GetUserByUsername(mUser)
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", user.Username, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.LoginFailur, nil)
		return
	}

	if !has { // 用户名不正确
		supports.Unauthorized(ctx, supports.UsernameFailur, nil)
		return
	}

	ckPassword := utils.CheckPWD(user.Password, mUser.Password)
	if !ckPassword {
		supports.Unauthorized(ctx, supports.PasswordFailur, nil)
		return
	}

	token, err := jwts.GenerateToken(mUser)
	golog.Infof("用户[%s], 登录生成token [%s]", mUser.Username, token)
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录，生成token出错。%s", user.Username, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.TokenCreateFailur, nil)
		return
	}

	supports.Ok(ctx, supports.LoginSuccess, vo.TansformUserVO(token, mUser))
	return
}

// 用户报表
func UserTable(ctx iris.Context) {
	page, err := supports.NewPagination(ctx)
	if err != nil {
		ctx.Application().Logger().Errorf("查询用户列表参数解析错误. %s", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	users, total, err := models.GetPaginationUsers(page)
	if err != nil {
		ctx.Application().Logger().Errorf("查询用户列表错误. %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	ctx.JSON(vo.BootstrapTableVO{
		Total: total,
		Rows:  vo.TansformUserVOList(users...),
	})
}

func UpdateUser(ctx iris.Context) {
	user := new(models.User)
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.Application().Logger().Errorf("更新用户[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.OptionFailur, nil)
		return
	}
	effect, err := models.UpdateUserById(user)
	if err != nil {
		ctx.Application().Logger().Errorf("更新用户[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, effect)
}

// 删除用户
func DeleteUsers(ctx iris.Context, uids string) {
	uidList := strings.Split(uids, ",")
	if len(uidList) == 0 {
		ctx.Application().Logger().Error("删除用户错误, 参数不对.")
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	dUids := make([]int64, 0)
	for _, v := range uidList {
		if v == "" {
			continue
		}
		uid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ctx.Application().Logger().Errorf("删除用户错误, %v", err)
			supports.Error(ctx, iris.StatusInternalServerError, supports.ParseParamsFailur, nil)
			return
		}
		dUids = append(dUids, uid)
	}

	effect, err := models.DeleteByUsers(dUids)
	if err != nil {
		ctx.Application().Logger().Errorf("删除用户错误, %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.DeleteUsersFailur, nil)
		return
	}
	supports.Ok(ctx, supports.DeleteUsersSuccess, effect)
}
