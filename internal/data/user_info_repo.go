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
	repo.log.Debug(" Save result  :", userInfo) // 返回 error

	if userInfo.Id > 0 { //更新	userInfo := biz.UserInfo{}
		var oldUserInfo biz.UserInfo
		result1 := repo.data.db.First(&oldUserInfo, "id = ? ", userInfo.Id) // 通过id进行数据查询
		repo.log.Debug("result.Error :", result1.Error)                     // 返回 error
		repo.log.Debug("save userInfo :", userInfo)
		if result1.Error != nil { // 有错误返回
			return userInfo, result1.Error
		} else {
			oldUserInfo.UserName = userInfo.UserName
			oldUserInfo.Password = userInfo.Password
			oldUserInfo.Age = userInfo.Age
			oldUserInfo.Phone = userInfo.Phone
			oldUserInfo.Address = userInfo.Address

			result := repo.data.db.Save(&oldUserInfo)      // 通过数据的指针来创建
			repo.log.Debug("result.Error :", result.Error) // 返回 error
			repo.log.Debug("save userInfo :", userInfo)
			return userInfo, result1.Error
		}

	} else { // 创建
		result := repo.data.db.Create(userInfo)        // 通过数据的指针来创建
		repo.log.Debug("result.Error :", result.Error) // 返回 error
		repo.log.Debug("save userInfo :", userInfo)
		return userInfo, result.Error
	}
}

func (repo *userInfoRepo) Update(ctx context.Context, userInfo *biz.UserInfo) error {
	repo.log.Debug("Update   :", userInfo) // 返回

	return nil
}

func (repo *userInfoRepo) Delete(ctx context.Context, id int64) error {
	repo.log.Debug("Delete  By Id :", id)                // 返回
	result := repo.data.db.Delete(&biz.UserInfo{Id: id}) // 通过id删除数据
	repo.log.Debug("result.Error :", result.Error)       // 返回 error
	return result.Error
}

func (repo *userInfoRepo) FindByID(ctx context.Context, id int64) (*biz.UserInfo, error) {
	repo.log.Debug("FindByID   :", id) // 返回
	userInfo := biz.UserInfo{}

	result := repo.data.db.First(&userInfo, "id = ? ", id) // 通过id进行数据查询
	repo.log.Debug("result.Error :", result.Error)         // 返回 error
	return &userInfo, result.Error
}

func (repo *userInfoRepo) FindAll(ctx context.Context) ([]*biz.UserInfo, error) {
	repo.log.Debug("FindAll   :") //
	var userInfoList []*biz.UserInfo

	result := repo.data.db.Find(&userInfoList)     // 通过数据查询
	repo.log.Debug("result.Error :", result.Error) // 返回 error
	return userInfoList, result.Error
}
