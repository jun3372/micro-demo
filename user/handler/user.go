package handler

import (
	"context"

	"github.com/micro/micro/v3/service/logger"

	proto "user/proto"
	"user/server"
)

type User struct {
	srv server.IUserServer
}

func NewUser() *User {
	return &User{
		srv: server.NewServer(),
	}
}

func (u User) Login(ctx context.Context, request *proto.LoginRequest, response *proto.LoginResponse) (err error) {
	logger.Info("Received User.Login request")
	// rsp, err := u.srv.Login(request)
	// if err != nil {
	// 	return err
	// }

	response.Err = 0
	response.Msg = "rsp.Msg"
	response.Data = &proto.JwtToken{Token: ""}
	return err
}

func (u User) Signup(ctx context.Context, request *proto.SignupRequest, response *proto.SignupResponse) error {
	panic("implement me")
}

// Call is a single request handler called via client.Call or the generated client code
// func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
// 	log.Info("Received User.Call request")
// 	rsp.Msg = "Hello " + req.Name
// 	return nil
// }
