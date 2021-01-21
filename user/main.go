package main

import (
	"user/handler"
	pb "user/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("users"),
		service.Version("latest"),
	)

	// Register handler
	_ = pb.RegisterUserHandler(srv.Server(), handler.NewUser())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
