package parse

import (
	"io/ioutil"

	"github.com/kataras/golog"

	"gopkg.in/yaml.v2"
)

var (
	DBConfig DB
)

func init() {
	golog.Info("@@@ Init db conf")
	ParseDBSetting()
}

func ParseDBSetting() {
	data, err := ioutil.ReadFile("conf/db.yml")
	//golog.Infof("%s", data)
	if err != nil {
		golog.Fatalf("@@@ %s", err)
	}
	err = yaml.Unmarshal(data, &DBConfig)
	if err != nil {
		golog.Fatalf("@@@ Unmarshal db config error!! %s", err)
	}
	//golog.Info(DBConfig)
}

type DB struct {
	Master DBConfigInfo
	Slave  DBConfigInfo
}

type DBConfigInfo struct {
	Dialect         string `yaml:"dialect"`
	URL             string `yaml:"url"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Database        string `yaml:"database"`
	Port            string `yaml:"port"`
	Host            string `yaml:"host"`
	Charset         string `yaml:"charset"`
	ParseTime       bool   `yaml:"parseTime"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int64  `yaml:"connMaxLifetime: 10"`
	Sslmode         string `yaml:"sslmode"`
}
