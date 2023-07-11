//go:build wireinject
// +build wireinject

package main

import (
	"costa92/gin-wire/config"
	"costa92/gin-wire/controllers"
	"costa92/gin-wire/dbs"
	"costa92/gin-wire/internal/app"
	"costa92/gin-wire/routers"
	"costa92/gin-wire/services"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func wireApp(*config.Configuration, *zap.Logger) (*app.App, func(), error) {
	panic(
		wire.Build(
			dbs.ProviderSet,
			services.ProviderServiceSet,
			controllers.ProviderSet,
			routers.ProviderSet,
			app.NewHttpServer,
			app.NewApp,
		),
	)
}
