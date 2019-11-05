package sys

import (
	"go-iris/middleware/casbins"

	"github.com/kataras/golog"
)

var (
	// 定义系统初始的角色
	Components = [][]string{
		{"admin", "/admin*", "GET|POST|DELETE|PUT", ".*", "角色管理"},
		{"demo", "/demo*", "GET|POST|DELETE|PUT", ".*", "demo角色"},
	}
)

// 创建系统默认角色
func CreateSystemRole() bool {
	e := casbins.GetEnforcer()

	for _, v := range Components {
		p := e.GetFilteredPolicy(0, v[0])
		if len(p) == 0 {
			if ok := e.AddPolicy(v); !ok {
				golog.Fatalf("初始化角色[%s]权限失败。", v)
			}
		}
	}
	return true
}
