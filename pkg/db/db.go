package db

import (
	"fmt"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/micro/micro/v3/service/config/client"
	"github.com/micro/micro/v3/service/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	sqlDB *gorm.DB
	once  sync.Once
)

func Init(namespace, project string) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		var cfg Config
		if cfg, err = InitConfig(namespace, project); err != nil {
			return
		}

		fmt.Println("cfg=", cfg, "Dialect", cfg.Dialect, "link=", cfg.Link())
		// _db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
		sqlDB, err = gorm.Open(mysql.Open(cfg.Link()), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",   // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		})

		// 全局禁用表名复数
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		// sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		// sqlDB.SetMaxOpenConns(100)
	})

	return sqlDB, err
}

func Get() *gorm.DB {
	return sqlDB
}

func InitConfig(namespace, project string) (Config, error) {
	// return Config{
	// 	Dialect:  "mysql",
	// 	Host:     "54.223.118.13",
	// 	Port:     3306,
	// 	User:     "root",
	// 	Password: "fe22a0fb4a94efae",
	// 	Dbname:   "demo",
	// 	Charset:  "utf8",
	// }, nil

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

		// 实例化配置
		config   = client.NewConfig(namespace)
		val, err = config.Get(project)
	)

	if err != nil {
		logger.Warn("获取数据库配置失败: err=", err)
		return def, err
	}

	if val.String("") == "" {
		logger.Warn("获取数据库配置内容失败: context=", val)
		return def, err
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.Unmarshal(val.Bytes(), &cfg); err != nil {
		logger.Warn("解密配置失败:err=", err)
	}
	return def, err
}

func (c Config) Link() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Dbname, c.Charset)
}
