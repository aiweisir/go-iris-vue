package routes

import (
	"go-iris/middleware/casbins"
	"go-iris/middleware/jwts"
	"go-iris/utils"
	"go-iris/web/models"
	"go-iris/web/services"
	"go-iris/web/supports"
	"go-iris/web/supports/vo"
	"net/http"
	"strconv"

	"github.com/kataras/iris"
)

func Registe(ctx iris.Context, u services.UserService) {
	user := new(models.User)
	ctx.ReadJSON(&user)

	err := u.DoRegiste(user)
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Username, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.Registe_failur, nil)
	} else {
		supports.Ok_(ctx, supports.Registe_success)
	}
}

func Login(ctx iris.Context, u services.UserService) {
	user := new(models.User)
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.Login_failur, nil)
		return
	}

	mUser := new(models.User)
	mUser.Username = user.Username
	has, err := u.DoLogin(mUser)
	//golog.Error(mUser)
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", user.Username, err.Error())
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

	token, err := jwts.GenerateToken(mUser);
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录，生成token出错。%s", user.Username, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.Token_create_failur, nil)
		return
	}

	supports.Ok(ctx, supports.Login_success, &vo.UserVO{
		Username:   mUser.Username,
		Appid:      mUser.Appid,
		Name:       mUser.Name,
		Phone:      mUser.Phone,
		Email:      mUser.Email,
		Userface:   mUser.Userface,
		CreateTime: mUser.CreateTime,
		Token:      token,
	})
	return
}

// 添加角色
func AddRole(ctx iris.Context) {
	roleDef := new(supports.RoleDefine)
	if err := ctx.ReadJSON(roleDef); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, nil)
	}
	if roleDef.Obj == "" {
		roleDef.Obj = "/demo/*"
	}
	if roleDef.Act == "" {
		roleDef.Act = "*"
	}
	if roleDef.Suf == "" {
		roleDef.Act = ".*"
	}

	e := casbins.GetEnforcer()
	ok := e.AddPolicy(utils.FmtRolePrefix(roleDef.Sub), roleDef.Obj, roleDef.Act, roleDef.Suf)
	if !ok {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, nil)
	}
	supports.Ok_(ctx, supports.Option_success)
}

// 修改角色的权限

// 给用户指定角色
func AddPermissions(ctx iris.Context) {
	groupDef := new(supports.GroupDefine)
	if err := ctx.ReadJSON(groupDef); err != nil {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, err.Error())
		return
	}

	var ok bool = true
	e := casbins.GetEnforcer()
	for _, v := range groupDef.Sub {
		if !e.AddGroupingPolicy(strconv.FormatInt(groupDef.Uid, 10), v) {
			ok = false
		}
	}

	if !ok {
		supports.Error(ctx, http.StatusInternalServerError, supports.Option_failur, nil)
	}
	supports.Ok_(ctx, supports.Option_success)
}
