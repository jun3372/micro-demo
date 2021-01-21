package db

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	jsoniter "github.com/json-iterator/go"
	"github.com/micro/micro/v3/service/config"
)

type Config struct {
	Dialect  string `json:"dialect"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Charset  string `json:"charset"`
}

var (
	_db  *gorm.DB
	once sync.Once
)

func Init(project string) (db *gorm.DB, err error) {
	once.Do(func() {
		var cfg Config
		if cfg, err = InitConfig(project); err != nil {
			return
		}

		// _db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
		_db, err = gorm.Open(cfg.Dialect, cfg.Link())
		_db.SingularTable(false)
	})

	return _db, err
}

func Get() *gorm.DB {
	return _db
}

func InitConfig(project string) (Config, error) {
	var (
		cfg Config
		def = Config{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Password: "root",
			Dbname:   "demo",
			Charset:  "utf8",
		}
		val, err = config.Get(project)
	)
	if err != nil {
		return def, err
	}

	context := val.String("")
	if context == "" {
		return def, err
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal([]byte(context), &cfg)
	return def, err
}

func (c Config) Link() string {
	return fmt.Sprintf("%s:%s@%s:%d/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Dbname, c.Charset)
}
