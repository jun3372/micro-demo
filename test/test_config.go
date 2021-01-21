package main

import (
	"fmt"

	"github.com/jun3372/micro-demo/pkg/db"
)

func main() {
	TestDB()
}

func TestDB() {
	_db, err := db.Init("db_member")
	if err != nil {
		panic(err)
	}

	fmt.Println("db=", _db)
}
