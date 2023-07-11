package services

import (
	"costa92/gin-wire/config"
	"costa92/gin-wire/dbs"
	"github.com/google/wire"
)

var ProviderServiceSet = wire.NewSet(NewApiService)

type ApiService struct {
	Database *dbs.Database
}

func NewApiService(cfg *config.Configuration, database *dbs.Database) *ApiService {
	return &ApiService{}
}
