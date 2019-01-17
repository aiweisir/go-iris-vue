package supports

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

// 设置分页参数
func (p *Pagination) PageSetting() {
	if p.PageNumber < 1 {
		p.PageNumber = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 1
	}

	p.Start = (p.PageNumber - 1) * p.PageSize
	p.Limit = p.PageSize
}

func NewPageCondition() {

}
