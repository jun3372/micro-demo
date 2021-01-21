package main

import (
	"github.com/jun3372/micro-demo/pkg/db"

	"member/handler"
	"member/model"
	pb "member/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("member"),
		service.Version("latest"),
	)

	_db, err := db.Init("db_member")
	if err != nil {
		panic(err)
	}

	_db.AutoMigrate(&model.Member{})

	// Register handler
	_ = pb.RegisterMemberHandler(srv.Server(), handler.NewMember())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
