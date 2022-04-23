package data

import (
	"kratos-crud/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserInfoRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	// mysql数据库连接
	//log.Infof("data soucre : %v", conf.Data_Database)
	dbUrl := "demo:demo@tcp(127.0.0.1:3306)/demo" //conf.Data_Database.Source
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	d := &Data{
		db: db,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
	}, nil
}
