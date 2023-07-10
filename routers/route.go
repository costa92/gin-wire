package routers

import (
	"costa92/gin-wire/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProviderSet is router providers.
var ProviderSet = wire.NewSet(NewRouter)

func NewRouter(conf *config.Configuration) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	// api group
	setApiGroupRoutes(router)
	return router
}
