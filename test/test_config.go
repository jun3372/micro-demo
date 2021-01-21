package main

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/jun3372/micro-demo/pkg/db"
)

func main() {
	TestDB()
}

type Demo struct {
	gorm.Model
	Avatar   string `json:"avatar" gorm:"type:varchar(100);"`
	Nickname string `json:"nickname" gorm:"type:varchar(100);"`
	Username string `json:"username" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password" gorm:"type:varchar(256)"`
}

func TestDB() {
	_db, err := db.Init("db_member")
	if err != nil {
		panic(err)
	}

	fmt.Println("db=", _db)
	if err = _db.AutoMigrate(&Demo{}); err != nil {
		panic(err)
	}
	fmt.Println("ok")
}
