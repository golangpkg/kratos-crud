package data

import (
	"context"
	"kratos-crud/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userInfoRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserInfoRepo .
func NewUserInfoRepo(data *Data, logger log.Logger) biz.UserInfoRepoIf {
	return &userInfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *userInfoRepo) Save(ctx context.Context, userInfo *biz.UserInfo) (*biz.UserInfo, error) {
	result := repo.data.db.Create(userInfo)        // 通过数据的指针来创建
	repo.log.Debug("result.Error :", result.Error) // 返回 error
	repo.log.Debug("save userInfo :", userInfo)
	return userInfo, nil
}

func (r *userInfoRepo) Update(ctx context.Context, g *biz.UserInfo) error {
	return nil
}

func (r *userInfoRepo) Delete(context.Context, int64) error {
	return nil
}

func (r *userInfoRepo) FindByID(context.Context, int64) (*biz.UserInfo, error) {
	return nil, nil
}

func (r *userInfoRepo) FindAll(context.Context) ([]*biz.UserInfo, error) {
	return nil, nil
}
