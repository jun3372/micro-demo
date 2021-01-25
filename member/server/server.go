package server

import (
	"fmt"

	"github.com/jun3372/micro-demo/pkg/db"
	"github.com/micro/micro/v3/service/logger"

	"member/model"
	pb "member/proto"
)

type IMemberServer interface {
	Login(req *pb.LoginRequest, rsp *pb.LoginResponse) error
	Signup(req *pb.SignupRequest, rsp *pb.SignupResponse) error
}

type MemberServer struct {
}

func (m MemberServer) Login(req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	var user model.Member
	first := db.Get().Where("`username`=?", req.Username).First(&user)
	if first.Error != nil {
		logger.Warnf("获取用户信息失败: %v", first.Error)
		rsp.Err = 1
		rsp.Msg = first.Error.Error()
		return nil
	}

	// 加密用户登录密码
	password := model.Member{Password: req.Password}.EncodePassword()
	if password != user.Password {
		rsp.Err = 1
		rsp.Msg = fmt.Sprint("登录密码错误")
		return nil
	}

	if token, err := user.MakeToken(); err != nil {
		logger.Warnf("生成用户提token失败: %v", err)
	} else {
		rsp.Data.Token = token
	}
	return nil
}

func (m MemberServer) Signup(req *pb.SignupRequest, rsp *pb.SignupResponse) error {
	var count int64
	db.Get().Model(&model.Member{}).Where("`username`=?", req.Username).Count(&count)
	if count > 0 {
		rsp.Err = 1
		rsp.Msg = fmt.Sprintf("已经存在该用户: username=%s", req.Username)
		logger.Warn(rsp.Msg)
		return nil
	}

	user := model.Member{
		Avatar:   req.GetAvatar(),
		Nickname: req.GetNickname(),
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	result := db.Get().Create(&user)
	if result.Error != nil {
		rsp.Err = 1
		rsp.Msg = result.Error.Error()
		logger.Warnf("注册用户失败: %v", result.Error)
		return nil
	}

	if token, err := user.MakeToken(); err != nil {
		logger.Warnf("生成用户提token失败: %v", err)
		rsp.Err = 1
		rsp.Msg = fmt.Sprintf("生成用户提token失败: %v", err)
		return nil
	} else {
		rsp.Data = &pb.JwtToken{Token: token}
	}

	return nil
}

func NewMemberServer() *MemberServer {
	return &MemberServer{}
}
