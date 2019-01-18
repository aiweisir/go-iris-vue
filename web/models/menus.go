package models

import (
	"fmt"
	"go-iris/web/db"
	"time"
)

/** gov doc
http://www.xorm.io/docs/
 */

type (
	// 菜单表
	Menu struct {
		Id          int64     `xorm:"pk autoincr INT(10) notnull" json:"id"`
		Path        string    `xorm:"varchar(64) notnull" json:"path"`
		Url         string    `xorm:"varchar(64) notnull" json:"url"`
		Modular     string    `xorm:"varchar(64) notnull" json:"modular"`
		Component   string    `xorm:"varchar(64) notnull" json:"component"`
		Name        string    `xorm:"varchar(64) notnull" json:"name"`
		Icon        string    `xorm:"varchar(64) notnull" json:"icon"`
		KeepAlive   string    `xorm:"varchar(64) notnull" json:"keepAlive"`
		RequireAuth string    `xorm:"varchar(64) notnull" json:"requireAuth"`
		ParentId    string    `xorm:"INT(10) notnull" json:"parentId"`
		Enabled     string    `xorm:"tinyint(1) notnull" json:"enabled"`
		CreateTime  time.Time `json:"createTime"`
		UpdateTime  time.Time `json:"updateTime"`

		Children []Children `json:"children"`
	}

	// 子菜单
	Children struct {
		Id2          int64  `xorm:"pk autoincr INT(10) notnull" json:"id"`
		Path2        string `xorm:"varchar(64) notnull" json:"path"`
		Modular2     string `xorm:"varchar(64) notnull" json:"modular"`
		Component2   string `xorm:"varchar(64) notnull" json:"component"`
		Name2        string `xorm:"varchar(64) notnull" json:"name"`
		Icon2        string `xorm:"varchar(64) notnull" json:"icon"`
		KeepAlive2   string `xorm:"varchar(64) notnull" json:"keepAlive"`
		RequireAuth2 string `xorm:"varchar(64) notnull" json:"requireAuth"`
		ParentId2    string `xorm:"INT(10) notnull" json:"parentId"`
		Enabled2     string `xorm:"tinyint(1) notnull" json:"enabled"`
	}

	// 菜单树
	MenuTreeGroup struct {
		Menu     `xorm:"extends"`
		Children `xorm:"extends"`
	}
)

// 获取用户的菜单列表
func DynamicMenuTree(uid int64) []Menu {
	e := db.MasterEngine()
	sql := fmt.Sprintf(`
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
				SELECT v1 FROM casbin_rule WHERE v0=%d
			)
		)
)
AND m2.enabled=true order by m1.id, m2.id
`, uid)

	menuTree := make([]MenuTreeGroup, 0)
	e.SQL(sql).Find(&menuTree)

	menus := make([]Menu, 0) // 菜单树

	if len(menuTree) > 0 {
		parentMenu := menuTree[0].Menu // 父级菜单
		childMenus := make([]Children, 0) // 所有的子菜单
		for _, v := range menuTree {
			childMenus = append(childMenus, v.Children)
		}
		parentMenu.Children = childMenus

		menus = append(menus, parentMenu)
	}
	return menus
}
