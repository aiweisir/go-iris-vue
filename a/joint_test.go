package a

import (
	"go-iris/web/models"
	"log"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
)

var (

)

func en() (*xorm.Engine) {
	url := "root:root@tcp(127.0.0.1:3306)/casbin?charset=utf8"
	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		log.Fatal(err.Error())
	}
	return engine
}

func TestRole(t *testing.T)  {
}

func TestJoint(t *testing.T)  {
	engine := en()

	join := make([]models.MenuTreeGroup, 0)

	sql := `
SELECT
	m1.id, m1.path, m1.modular, m1.component, m1.icon, m1.name, m1.require_auth,
	m2.id AS id2,
	m2.modular AS modular2,
	m2.component AS component2,
	m2.icon AS icon2,
	m2.keep_alive AS keep_alive2,
	m2.name AS name2,
	m2.path AS path2,
	m2.require_auth AS require_auth2
FROM menu m1, menu m2
WHERE m1.id = m2.parent_id
	AND m1.id != 1
	AND m2.id IN 
(
		SELECT rm.mid
		FROM role_menu rm WHERE rm.rid in
		(
			SELECT id FROM casbin_rule 
			WHERE 
			v2 <> 'ANY' AND 
			v0 in 
			(
				SELECT v1 FROM casbin_rule WHERE v0=90
			)
		)
)
AND m2.enabled=true order by m1.id, m2.id
`
	engine.SQL(sql).Find(&join)
	//t.Log(join)
	parent := join[0].Menu
	child := make([]models.Children, 0)
	for _, v := range join {
		child = append(child, v.Children)
	}

	parent.Children = child
	//t.Log("p => ", parent)
	//for _, v := range child {
	//	t.Log("c => ", v)
	//}

	user := new(models.User)
	i, e := engine.Where("id = ?", 90).Count(user)
	golog.Info(i, e)
}
