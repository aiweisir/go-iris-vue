package parse

import (
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
	//data, err := ioutil.ReadFile("conf/db.yml")
	////golog.Infof("%s", data)
	//if err != nil {
	//	golog.Fatalf("@@@ %s", err)
	//}
	//err = yaml.Unmarshal(data, &DBConfig)
	//if err != nil {
	//	golog.Fatalf("@@@ Unmarshal db config error!! %s", err)
	//}

	dbData, err := Asset("conf/db.yml")
	if err != nil {
		golog.Fatalf("Error. %s", err)
	}
	if err = yaml.Unmarshal(dbData, &DBConfig); err != nil {
		golog.Fatalf("Error. %s", err)
	}

	//golog.Info(DBConfig)
}

type DB struct {
	Master DBConfigInfo
	Slave  DBConfigInfo
}

type DBConfigInfo struct {
	Dialect  string `yaml:"dialect"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
	ShowSql  bool   `yaml:"showSql"`
	LogLevel string `yaml:"logLevel"`
	MaxIdleConns int `yaml:"maxIdleConns"`
	MaxOpenConns int `yaml:"maxOpenConns"`

	//ParseTime       bool   `yaml:"parseTime"`
	//MaxIdleConns    int    `yaml:"maxIdleConns"`
	//MaxOpenConns    int    `yaml:"maxOpenConns"`
	//ConnMaxLifetime int64  `yaml:"connMaxLifetime: 10"`
	//Sslmode         string `yaml:"sslmode"`
}
