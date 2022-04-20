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
	Age      int
	Phone    string
	Address  string
}

// UserInfoRepo is a Greater repo.
type UserInfoRepo interface {
	Save(context.Context, *UserInfo) (*UserInfo, error)
	Delete(context.Context, int64) (*UserInfo, error)
	Update(context.Context, *UserInfo) error
	FindByID(context.Context, int64) (*UserInfo, error)
	ListAll(context.Context) ([]*UserInfo, error)
}

// UserInfoUsecase is a UserInfo usecase.
type UserInfoUsecase struct {
	repo UserInfoRepo
	log  *log.Helper
}

// NewUserInfoUsecase new a UserInfo usecase.
func NewUserInfoUsecase(repo UserInfoRepo, logger log.Logger) *UserInfoUsecase {
	return &UserInfoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUserInfo creates a UserInfo, and returns the new UserInfo.
func (uc *UserInfoUsecase) CreateUserInfo(ctx context.Context, g *UserInfo) (*UserInfo, error) {
	uc.log.WithContext(ctx).Infof("CreateUserInfo: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
