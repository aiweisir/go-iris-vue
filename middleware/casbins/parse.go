package casbins

import (
	"github.com/kataras/golog"
)

/**
用于解析权限
 */

// 通过uid获取用户的所有资源
func GetAllResourcesByUID(uid string) map[string]interface{} {
	allRes := make(map[string]interface{})

	e := GetEnforcer()

	myRes := e.GetPermissionsForUser(uid)
	golog.Infof("myRes=> %s", myRes)

	// 获取用户的隐形角色
	implicitRoles := e.GetImplicitRolesForUser(uid)
	for _, v := range implicitRoles{
		// 查询用户隐形角色的资源权限
		subRes := e.GetPermissionsForUser(v)
		golog.Infof("-------------------------------------------------")
		golog.Infof("subRes[%s], len(res)=> %d", v, len(subRes))
		golog.Infof("subRes[%s], res=> %s", v, subRes)
		golog.Infof("-------------------------------------------------")
		allRes[v] = subRes
	}

	allRes["myRes"] = myRes
	return allRes
}

// 通过uid获取用户的所有角色
func GetAllRoleByUID(uid string) []string {
	e := GetEnforcer()
	roles := e.GetImplicitRolesForUser(uid)
	golog.Infof("roles=> %s", roles)
	return roles
}
