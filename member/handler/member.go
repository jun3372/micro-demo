package handler

import (
	"context"

	pb "member/proto"
	"member/server"
)

type Member struct {
	srv server.IMemberServer
}

func NewMember() *Member {
	return &Member{
		srv: server.NewMemberServer(),
	}
}

func (m Member) Login(ctx context.Context, request *pb.LoginRequest, response *pb.LoginResponse) error {
	return m.srv.Login(request, response)
}

func (m Member) Signup(ctx context.Context, request *pb.SignupRequest, response *pb.SignupResponse) error {
	return m.srv.Signup(request, response)
}
