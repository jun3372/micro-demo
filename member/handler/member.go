package handler

import (
	"context"

	pb "member/proto"
)

type Member struct{}

func NewMember() *Member {
	return &Member{}
}

func (m Member) Login(ctx context.Context, request *pb.LoginRequest, response *pb.LoginResponse) error {
	response.Msg = "err:username=" + request.GetUsername()
	response.Data = &pb.JwtToken{Token: "password:" + request.GetPassword()}
	return nil
}

func (m Member) Signup(ctx context.Context, request *pb.SignupRequest, response *pb.SignupResponse) error {
	panic("implement me")
}

//
// // Call is a single request handler called via client.Call or the generated client code
// func (e *Member) Call(ctx context.Context, req *member.Request, rsp *member.Response) error {
// 	log.Info("Received Member.Call request")
// 	rsp.Msg = "Hello " + req.Name
// 	return nil
// }
