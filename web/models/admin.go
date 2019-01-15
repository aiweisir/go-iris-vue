package models

import (
	"fmt"
	"go-iris/web/db"
	"time"
)

/**
http://www.xorm.io/docs/
  */

// 菜单表
type Menu struct {
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
}

// 角色-菜单关联表
type RoleMenu struct {
	Id  int64 `xorm:"pk autoincr INT(10) notnull" json:"id"`
	Rid int64 `xorm:"pk autoincr INT(10) notnull" json:"rid"`
	Mid int64 `xorm:"pk autoincr INT(10) notnull" json:"mid"`
}

// 菜单树
type MenuTree struct {
	Id          int64  `xorm:"pk autoincr INT(10) notnull" json:"id"`
	Path        string `xorm:"varchar(64) notnull" json:"path"`
	Modular     string `xorm:"varchar(64) notnull" json:"modular"`
	Component   string `xorm:"varchar(64) notnull" json:"component"`
	Name        string `xorm:"varchar(64) notnull" json:"name"`
	Icon        string `xorm:"varchar(64) notnull" json:"icon"`
	KeepAlive   string `xorm:"varchar(64) notnull" json:"keepAlive"`
	RequireAuth string `xorm:"varchar(64) notnull" json:"requireAuth"`
	ParentId    string `xorm:"INT(10) notnull" json:"parentId"`
	Enabled     string `xorm:"tinyint(1) notnull" json:"enabled"`

	Id2          int64  `xorm:"pk autoincr INT(10) notnull" json:"id2"`
	Path2        string `xorm:"varchar(64) notnull" json:"path2"`
	Modular2     string `xorm:"varchar(64) notnull" json:"modular2"`
	Component2   string `xorm:"varchar(64) notnull" json:"component2"`
	Name2        string `xorm:"varchar(64) notnull" json:"name2"`
	Icon2        string `xorm:"varchar(64) notnull" json:"icon2"`
	KeepAlive2   string `xorm:"varchar(64) notnull" json:"keepAlive2"`
	RequireAuth2 string `xorm:"varchar(64) notnull" json:"requireAuth2"`
	ParentId2    string `xorm:"INT(10) notnull" json:"parentId2"`
	Enabled2     string `xorm:"tinyint(1) notnull" json:"enabled2"`
}

func DynamicMenuTree(uid int64) *[]MenuTree {
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
			act <> 'ANY' AND 
			sub in 
			(
				SELECT obj FROM casbin_rule WHERE sub=%d
			)
		)
)
AND m2.enabled=true order by m1.id, m2.id
`, uid)

	join := make([]MenuTree, 0)
	e := db.MasterEngine()
	e.SQL(sql).Find(&join)

	return &join
}


