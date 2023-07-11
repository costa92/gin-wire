package dbs

import (
	"costa92/gin-wire/config"
	"costa92/gin-wire/pkg"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewDatabase)

type Database struct {
	Master *gorm.DB
}

func NewDatabase(cfg *config.Configuration) (*Database, error) {
	masterDB, err := pkg.InitGormV2(&cfg.MasterDB)
	if err != nil {
		return nil, err
	}
	return &Database{
		Master: masterDB,
	}, err
}
