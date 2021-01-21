package server

import (
	"fmt"

	user "user/proto"
)

type IUserServer interface {
	Login(req *user.LoginRequest) (rsp *user.LoginResponse, err error)
	Signup(req *user.SignupRequest) (rsp *user.SignupResponse, err error)
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Login(req *user.LoginRequest) (rsp *user.LoginResponse, err error) {
	rsp = &user.LoginResponse{
		Err:  0,
		Msg:  "field",
		Data: &user.JwtToken{Token: fmt.Sprintf("token:%s:%s", req.Username, req.Password)},
	}

	return rsp, err
}

func (s Server) Signup(req *user.SignupRequest) (rsp *user.SignupResponse, err error) {
	panic("implement me")
}
