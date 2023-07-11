package controllers

import (
	"costa92/gin-wire/config"
	"costa92/gin-wire/services"
	"github.com/google/wire"
)

// ProviderSet is router providers.
var ProviderSet = wire.NewSet(ProvideTodoApiController)

type ApiController struct {
	TodoAPI     *TodoAPI
	ApiServices *services.ApiService
}

func ProvideTodoApiController(cfg *config.Configuration, services *services.ApiService) *ApiController {
	wire.Build(TodoAPISet, services)
	return &ApiController{}
}
