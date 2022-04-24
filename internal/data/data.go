package data

import (
	"kratos-crud/internal/conf"
	"time"

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
	dbUrl := c.Database.Source
	log.Infof("data soucre : %v", dbUrl)
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	sqlDB, err := db.DB()
	if err != nil {
		log.Error("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(int(c.Database.MaxIdleConns))
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(int(c.Database.MaxOpenConns))
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(c.Database.ConnMaxLifetime))
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

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
