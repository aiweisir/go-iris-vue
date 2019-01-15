package parse

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"gopkg.in/yaml.v2"
)

var (
	// conf strut
	C iris.Configuration

	// 解析app.yml中的Other项
	O Other
	// app.conf配置项key定义
	ignoreURLs string = "IgnoreURLs"
	jwtTimeout string = "JWTTimeout"
	logLevel   string = "LogLevel"
	secret     string = "Secret"
)

type (
	Other struct {
		IgnoreURLs []string
		JWTTimeout int64
		LogLevel   string
		Secret     string
	}
)

func AppOtherParse() {
	golog.Info("@@@ Init app conf")
	//c := iris.YAML("conf/app.yml")

	appData, err := Asset("conf/app.yml")
	if err != nil {
		golog.Fatalf("Error. %s", err)
	}
	c := iris.DefaultConfiguration()
	if err = yaml.Unmarshal(appData, &c); err != nil {
		golog.Fatalf("Error. %s", err)
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

	O.LogLevel = c.GetOther()[logLevel].(string)

	O.Secret = c.GetOther()[secret].(string)

}
