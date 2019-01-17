package routes

import (
	"go-iris/middleware/jwts"
	"go-iris/utils"
	"go-iris/web/models"
	"go-iris/web/supports"
	"go-iris/web/supports/vo"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func Registe(ctx iris.Context) {
	user := new(models.User)
	ctx.ReadJSON(&user)

	user.CreateTime = time.Now()
	user.Password = utils.AESEncrypt([]byte(user.Password))

	effect, err := models.CreateUser(user)
	//err := u.DoRegiste(user)
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
		supports.Error(ctx, iris.StatusInternalServerError, supports.LoginFailur, nil)
		return
	}

	mUser := new(models.User)
	mUser.Username = user.Username
	has, err := models.GetUserByUsername(mUser)
	//has, err := u.DoLogin(mUser)
	//golog.Error(mUser)
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

	token, err := jwts.GenerateToken(mUser);
	golog.Infof("用户[%s], 登录生成token [%s]", mUser.Username, token)
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录，生成token出错。%s", user.Username, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.TokenCreateFailur, nil)
		return
	}

	supports.Ok(ctx, supports.LoginSuccess, vo.TansformUserVO(token, mUser))
	return
}

func UserTable(ctx iris.Context) {
	pageNumber, err1 := ctx.URLParamInt("pageNumber")
	pageSize, err2 := ctx.URLParamInt("pageSize")
	sortName := ctx.URLParam("sortName")
	sortOrder := ctx.URLParam("sortOrder")
	golog.Infof("pageNumber=%d, pageSize=%d, sortName=%s, sortOrder=%s", pageNumber, pageSize, sortName, sortOrder)
	if err1 != nil || err2 != nil {
		ctx.Application().Logger().Errorf("查询用户列表参数解析错误. %s, %s", err1.Error(), err2.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.ParseParamsFailur, nil)
		return
	}

	page := supports.Pagination{
		PageNumber:pageNumber,
		PageSize:pageSize,
		SortName:sortName,
		SortOrder:sortOrder,
	}
	page.PageSetting()

	users, total, err := models.GetPaginationUsers(&page)
	if err != nil {
		ctx.Application().Logger().Errorf("查询用户列表错误. %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	ctx.JSON(vo.BootstrapTableVO{
		Total:total,
		Rows:vo.TansformUserVOList(users...),
	})
}

// 修改角色的权限


