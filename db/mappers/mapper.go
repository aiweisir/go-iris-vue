package mappers

import (
	"fmt"
	"reflect"

	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	adt := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e := casbin.NewEnforcer("conf/rbac_model.conf", adt)

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
