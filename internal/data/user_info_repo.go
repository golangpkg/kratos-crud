package data

import (
	"context"

	"kratos-crud/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserInfoRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserInfoRepo .
func NewUserInfoRepo(data *Data, logger log.Logger) *UserInfoRepo {
	return &UserInfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *UserInfoRepo) Save(ctx context.Context, userInfo *biz.UserInfo) (*biz.UserInfo, error) {
	result := repo.data.db.Create(userInfo)        // 通过数据的指针来创建
	repo.log.Debug("result.Error :", result.Error) // 返回 error
	repo.log.Debug("save userInfo :", userInfo)
	return userInfo, nil
}

func (r *UserInfoRepo) Update(ctx context.Context, g *biz.UserInfo) (*biz.UserInfo, error) {
	return g, nil
}

func (r *UserInfoRepo) FindByID(context.Context, int64) (*biz.UserInfo, error) {
	return nil, nil
}

func (r *UserInfoRepo) ListByHello(context.Context, string) ([]*biz.UserInfo, error) {
	return nil, nil
}

func (r *UserInfoRepo) ListAll(context.Context) ([]*biz.UserInfo, error) {
	return nil, nil
}
