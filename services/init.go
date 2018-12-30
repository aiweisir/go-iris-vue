package services

import (
	"casbin-demo/db/mappers"
	"casbin-demo/routes/dispatch"
	"github.com/kataras/golog"
)

// 统一注入service
func init() {
	golog.Info("@@@ init injection services")

	userService := NewUserService(mappers.NewUserMapper())

	dispatch.Register(
		userService,

		)
}
