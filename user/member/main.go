package main

import (
	"member/handler"
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

	// Register handler
	pb.RegisterMemberHandler(srv.Server(), new(handler.Member))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
