package api

import (
	"context"
	pb "entry_task/src/proto"
	"entry_task/src/server/tcpserver/code"
	"entry_task/src/server/tcpserver/service"

	"entry_task/src/tools"
	"fmt"
	log "github.com/cihub/seelog"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type UserServer struct {
}

func (server *UserServer) Login(ctx context.Context, in *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {

	// 根据用户名查询用户信息
	user, err := service.GetUserInfo(in.Username)
	if err != nil {
		log.Errorf("【rpc服务】获取用户信息发生异常，用户名:%s,密码：%s,异常信息：%s", in.Username, in.Password, err.Error())
		return &pb.UserLoginResponse{Code: code.CodeRPCSystemBusy, Msg: code.CodeMsg[code.CodeRPCSystemBusy]}, nil
	}
	if user.Username == "" {
		return &pb.UserLoginResponse{Code: code.CodeRPCUserNotExist, Msg: code.CodeMsg[code.CodeRPCUserNotExist]}, nil
	}
	// 校验密码，输入密码和用户的秘钥进行md5加密后是否与数据库中的密码一致
	log.Debugf("【rpc服务】加密前的密码：%s")
	passwordAfterMd5 := tools.GetMd5String(in.Password + user.SecretKey)
	log.Debugf("【rpc服务】加密后的密码：%s")
	if passwordAfterMd5 != user.Password {
		log.Debug("【rpc服务】密码不正确")
		return &pb.UserLoginResponse{Code: code.CodeRPCFailedToVerifyPassword, Msg: code.CodeMsg[code.CodeRPCFailedToVerifyPassword]}, nil
	}
	return &pb.UserLoginResponse{Code: code.CodeRPCSuccess,
	}, nil
}

func (server *UserServer) Logout(ctx context.Context, in *pb.CommonRequest) (*pb.CommonResponse, error) {
	return &pb.CommonResponse{Code: code.CodeRPCSuccess}, nil
}

func (server *UserServer) GetUserInfo(ctx context.Context, in *pb.CommonRequest) (*pb.UserInfoResponse, error) {
	user, e := service.GetUserInfo(in.Username)
	if e != nil {
		log.Errorf("【rpc服务】获取用户信息发生异常，用户名:%s,异常信息：%s", in.Username, e.Error())
		return &pb.UserInfoResponse{Code: code.CodeRPCSystemBusy, Msg: code.CodeMsg[code.CodeRPCSystemBusy]}, nil
	}

	if user.Username == "" {
		return &pb.UserInfoResponse{Code: code.CodeRPCUserNotExist, Msg: code.CodeMsg[code.CodeRPCUserNotExist]}, nil
	}
	return &pb.UserInfoResponse{Username: user.Username, Nickname: user.Nickname, ProfilePicture: user.ProfilePicture, Code: code.CodeRPCSuccess}, nil
}

func (server *UserServer) EditUserInfo(ctx context.Context, in *pb.EditUserInfoRequest) (*pb.CommonResponse, error) {
	// 修改用户信息
	service.UpdateUserInfo(in.Username, in.Nickname, in.ProfilePicture, "")
	return &pb.CommonResponse{Code: code.CodeRPCSuccess, Msg: code.CodeMsg[code.CodeRPCSuccess]}, nil
}

func Start() {
	address := fmt.Sprintf("%s:%s", tools.GetTcpConfig().Server.Ip, strconv.Itoa(tools.GetTcpConfig().Server.Port))
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf(" 【rpc服务】监听tcp端口发生异常，异常信息 ", err.Error())
		return
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServer{})
	err = server.Serve(listen)
	if err != nil {
		log.Errorf("【rpc】服务启动发生异常，异常信息：%s", err.Error())
	}
	log.Debugf("【rpc服务】成功启动，address:%s", address)
}
