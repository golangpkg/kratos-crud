package service

import (
	"context"

	v1 "kratos-crud/api/helloworld/v1"
	"kratos-crud/internal/biz"
)

// UserInfoService is a UserInfo service.
type UserInfoService struct {
	v1.UnimplementedUserInfoServer

	uc *biz.UserInfoUsecase
}

// NewUserInfoService new a UserInfo service.
func NewUserInfoService(uc *biz.UserInfoUsecase) *UserInfoService {
	return &UserInfoService{uc: uc}
}

// SayHello implements helloworld.UserInfoServer.
func (s *UserInfoService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateUserInfo(ctx, &biz.UserInfo{UserName: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.UserName}, nil
}
