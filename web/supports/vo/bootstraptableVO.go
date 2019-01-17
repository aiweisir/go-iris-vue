package vo

type BootstrapTableVO struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}
