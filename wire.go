//go:build wireinject
// +build wireinject

package main

import (
	"costa92/gin-wire/config"
	"costa92/gin-wire/routers"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// wireApp init application.
func wireApp(*config.Configuration, *zap.Logger) (*App, func(), error) {
	panic(
		wire.Build(
			routers.ProviderSet,
			newHttpServer,
			newApp,
		),
	)
}
