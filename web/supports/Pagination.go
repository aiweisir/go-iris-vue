package supports

import (
	"errors"

	"github.com/kataras/iris"
)

// bootstraptable 分页参数
type Pagination struct {
	PageNumber int //当前看的是第几页
	PageSize   int //每页显示多少条数据

	// 用于分页设置的参数
	Start int
	Limit int

	SortName  string //用于指定的排序
	SortOrder string // desc或asc

	// 时间范围
	StartDate string
	EndDate   string

	Uid int64 // 公用的特殊参数
}

func NewPagination(ctx iris.Context) (*Pagination, error) {
	pageNumber, err1 := ctx.URLParamInt("pageNumber")
	pageSize, err2 := ctx.URLParamInt("pageSize")
	sortName := ctx.URLParam("sortName")
	sortOrder := ctx.URLParam("sortOrder")
	if err1 != nil || err2 != nil {
		return nil, errors.New("请求的分页参数解析错误.")
	}

	page := Pagination{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		SortName:   sortName,
		SortOrder:  sortOrder,
	}
	page.pageSetting()
	return &page, nil
}

// 设置分页参数
func (p *Pagination) pageSetting() {
	if p.PageNumber < 1 {
		p.PageNumber = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 1
	}

	p.Start = (p.PageNumber - 1) * p.PageSize
	p.Limit = p.PageSize
}
