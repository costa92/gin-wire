//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"kart-cloud/kart/example/base-example/app"
)

func wireApp(*gin.Engine) (*app.App, func(), error) {
	panic(wire.Build(app.ProviderHttpSeverSet))
}
