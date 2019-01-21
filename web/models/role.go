package models

import (
	"go-iris/web/db"
	"go-iris/web/supports"
)

type CasbinRule struct {
	Id    int64  `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	PType string `xorm:"varchar(100) index" json:"p_type"`
	V0    string `xorm:"varchar(100) index" json:"v0"`
	V1    string `xorm:"varchar(100) index" json:"v1"`
	V2    string `xorm:"varchar(100) index" json:"v2"`
	V3    string `xorm:"varchar(100) index" json:"v3"`
	V4    string `xorm:"varchar(100) index" json:"v4"`
	V5    string `xorm:"varchar(100) index" json:"v5"`
}

func GetPaginationRoles(page *supports.Pagination) ([]*CasbinRule, int64, error) {
	e := db.MasterEngine()
	roleList := make([]*CasbinRule, 0)

	s := e.Where("p_type=?", "p").Limit(page.Limit, page.Start)
	if page.SortName != "" {
		switch page.SortOrder {
		case "asc":
			s.Asc(page.SortName)
		case "desc":
			s.Desc(page.SortName)
		}
	}
	count, err := s.FindAndCount(&roleList)

	return roleList, count, err
}

func UpdateRoleById(role *CasbinRule) (int64, error) {
	e := db.MasterEngine()
	return e.Id(role.Id).Update(role)
}

func DeleteByRoles(rids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	cr := new(CasbinRule)
	for _, v := range rids {
		i, err1 := e.Id(v).Delete(cr)
		effect += i
		err = err1
	}
	return
}
