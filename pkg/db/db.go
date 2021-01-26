package db

import (
	"fmt"
	"strings"
	"sync"

	"github.com/micro/micro/v3/service/config"
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
	Link     string `json:"link"`
}

var (
	sqlDB *gorm.DB
	once  sync.Once
)

func Init(project string) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		var cfg *Config
		if cfg, err = InitConfig(project); err != nil {
			return
		}

		// _db, err := gorm.Open("user:password@/dbname?charset=utf8&parseTime=True&loc=Local", &gorm.Config)
		sqlDB, err = gorm.Open(cfg.Director(), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",   // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		})
	})

	return sqlDB, err
}

func Get() *gorm.DB {
	return sqlDB
}

func InitConfig(project string) (*Config, error) {
	var (
		cfg *Config

		// 实例化配置
		val, err = config.Get(project)
	)

	if err != nil {
		logger.Fatalf("获取数据库配置失败: err=", err)
		return nil, err
	}

	if err = val.Scan(&cfg); err != nil {
		logger.Fatalf("转义配置结构体失败: err=", err)
		return nil, err
	}

	return cfg, err
}

func (c Config) GetLink() string {
	if c.Link != "" {
		return c.Link
	}

	var link = ""
	switch strings.ToUpper(c.Dialect) {
	case "MYSQL":
	default:
		link = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Dbname, c.Charset)
	}

	return link
}

// 获取连接驱动
func (c Config) Director() gorm.Dialector {
	switch strings.ToUpper(c.Dialect) {
	case "MYSQL":
	default:
		return mysql.Open(c.GetLink())
	}

	return mysql.Open(c.GetLink())
}
