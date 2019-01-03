package parse

import (
	"github.com/kataras/iris"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v2"
)

func init() {
	golog.Info("@@@ Init app conf")
	//c := iris.YAML("conf/app.yml")

	appData, err := Asset("conf/app.yml")
	if err != nil {
		golog.Fatalf("Get Asset of app-data.go error. %s", err)
	}
	c := iris.DefaultConfiguration()
	if err = yaml.Unmarshal(appData, &c); err != nil {
		golog.Fatalf("Unmarshal Asset of app-data.go error. %s", err)
	}
	C = c

	// 解析other的key
	iURLs := c.GetOther()[ignoreURLs].([]interface{})
	for _, v := range iURLs {
		O.IgnoreURLs = append(O.IgnoreURLs, v.(string))
	}

	jTimeout := c.GetOther()[jwtTimeout].(int)
	O.JWTTimeout = int64(jTimeout)
	//golog.Info(reflect.TypeOf(O.JWTTimeout))

	logLvl := c.GetOther()[logLevel].(string)
	O.LogLevel = logLvl
}


var (
	// conf strut
	C iris.Configuration

	// 解析app.yml中的Other项
	O Other
	// app.conf配置项key定义
	ignoreURLs string = "IgnoreURLs"
	jwtTimeout string = "JWTTimeout"
	logLevel string = "LogLevel"
)

type (
	Other struct {
		IgnoreURLs []string
		JWTTimeout int64
		LogLevel string
	}
)
