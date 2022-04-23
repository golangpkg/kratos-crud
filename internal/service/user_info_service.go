package service

import (
	"context"
	v1 "kratos-crud/api/helloworld/v1"
	"kratos-crud/internal/biz"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

// UserInfoService is a UserInfo service.
type UserInfoService struct {
	v1.UnimplementedUserInfoServiceServer
	uc  *biz.UserInfoUsecase
	log *log.Helper
}

// NewUserInfoService new a UserInfo service.
func NewUserInfoService(uc *biz.UserInfoUsecase, logger log.Logger) *UserInfoService {

	log := log.NewHelper(logger)

	return &UserInfoService{uc: uc, log: log}
}

// Save
func (s *UserInfoService) Save(ctx context.Context, in *v1.UserInfo) (*v1.CommReply, error) {
	s.log.Info(in.GetUserName())
	userInfo := biz.UserInfo{}
	userInfo.UserName = in.GetUserName()
	userInfo.Password = in.GetPassword()
	userInfo.Age = in.GetAge()
	userInfo.Phone = in.GetPhone()
	userInfo.Address = in.GetAddress()

	g, err := s.uc.CreateUserInfo(ctx, &userInfo)
	if err != nil {
		return nil, err
	}
	return &v1.CommReply{Message: "Hello " + g.UserName}, nil
}

// Delete
func (s *UserInfoService) Delete(ctx context.Context, in *v1.IdRequest) (*v1.CommReply, error) {
	err := s.uc.DeleteUserInfo(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.CommReply{Message: "Delete " + strconv.FormatInt(in.GetId(), 10)}, nil
}

// Get
func (s *UserInfoService) Get(ctx context.Context, in *v1.IdRequest) (*v1.UserInfoReply, error) {
	user, err := s.uc.FindUserInfoByID(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoReply{Message: "Get " + user.UserName}, nil
}

// List
func (s *UserInfoService) List(ctx context.Context, in *v1.ListRequest) (*v1.ListUserInfoReply, error) {
	userList, err := s.uc.FindAllUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	log.Info(userList)
	return &v1.ListUserInfoReply{Message: "List UserInfo"}, nil
}
