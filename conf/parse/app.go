package parse

import (
	"github.com/kataras/iris"
)

// 解析app.yml中的Other项
var (
	O Other

	ignoreURLs string = "IgnoreURLs"
)

type (
	Other struct {
		IgnoreURLs []string
	}
)

func InitOtherConfig(c *iris.Configuration) {
	iURLs := c.GetOther()[ignoreURLs].([]interface{})
	for _, v := range iURLs {
		O.IgnoreURLs = append(O.IgnoreURLs, v.(string))
	}

}
