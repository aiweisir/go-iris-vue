package a

import (
	"fmt"
	"reflect"

	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

func Enforce() *casbin.Enforcer {
	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbins".
	// If it doesn't exist, the adapter will create it automatically.
	adt := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	e := casbin.NewEnforcer("conf/rbac_model.conf", adt)
	return e
}

func a1()  {
	e := Enforce()
	// Load the policy from DB.
	e.LoadPolicy()


	e.AddGroupingPolicy("1", "alice")

	e.GetPolicy()
	// Check the permission.
	//check := e.Enforce(e.GetPolicy())
	fmt.Println(reflect.TypeOf(e.GetPolicy()))

	// Modify the policy.
	e.AddPolicy("alice", "data1", "read")
	// e.RemovePolicy(...)

	e.Enforce()
	//e.GetPermissionsForUser()

	// Save the policy back to DB.
	e.SavePolicy()
}

func a2()  {
	//e := Enforce()

}

func main()  {
	a2()
}
