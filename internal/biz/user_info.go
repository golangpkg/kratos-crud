package biz

import (
	"context"

	v1 "kratos-crud/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

/**
CREATE TABLE `user_info` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`user_name` varchar(225) NOT NULL,
`password` varchar(225) DEFAULT NULL,
`age` tinyint(4) DEFAULT NULL,
`phone` varchar(20) DEFAULT NULL,
`address` varchar(255) DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
*/
// UserInfo model.
type UserInfo struct {
	Id       int64
	UserName string
	Password string
	Age      uint32
	Phone    string
	Address  string
}

func (UserInfo) TableName() string {
	return "user_info"
}

// UserInfoRepo is a Greater repo.
type UserInfoRepoIf interface {
	Save(context.Context, *UserInfo) (*UserInfo, error)
	Delete(context.Context, int64) error
	Update(context.Context, *UserInfo) error
	FindByID(context.Context, int64) (*UserInfo, error)
	FindAll(context.Context) ([]*UserInfo, error)
}

// UserInfoUsecase is a UserInfo usecase.
type UserInfoUsecase struct {
	repo UserInfoRepoIf
	log  *log.Helper
}

// NewUserInfoUsecase new a UserInfo usecase.
func NewUserInfoUsecase(repo UserInfoRepoIf, logger log.Logger) *UserInfoUsecase {
	return &UserInfoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUserInfo
func (uc *UserInfoUsecase) CreateUserInfo(ctx context.Context, user *UserInfo) (*UserInfo, error) {
	uc.log.WithContext(ctx).Infof("CreateUserInfo: %v", user.UserName)
	return uc.repo.Save(ctx, user)
}

// DeleteUserInfo
func (uc *UserInfoUsecase) DeleteUserInfo(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteUserInfo: %v", id)
	return uc.repo.Delete(ctx, id)
}

// UpdateUserInfo
func (uc *UserInfoUsecase) UpdateUserInfo(ctx context.Context, user *UserInfo) error {
	uc.log.WithContext(ctx).Infof("UpdateUserInfo: %v", user)
	return uc.repo.Update(ctx, user)
}

// FindUserInfoByID
func (uc *UserInfoUsecase) FindUserInfoByID(ctx context.Context, id int64) (*UserInfo, error) {
	uc.log.WithContext(ctx).Infof("FindUserInfoByID: %v", id)
	return uc.repo.FindByID(ctx, id)
}

// FindAllUserInfo
func (uc *UserInfoUsecase) FindAllUserInfo(ctx context.Context) ([]*UserInfo, error) {
	uc.log.WithContext(ctx).Infof("FindAllUserInfo ")
	return uc.repo.FindAll(ctx)
}
