package server

import (
	pb "member/proto"
)

type IMemberServer interface {
	Login(response pb.LoginResponse)
}