package routers

import (
	"costa92/gin-wire/controllers"
	"github.com/gin-gonic/gin"
)

func setApiGroupRoutes(router *gin.Engine, apiCtx *controllers.ApiController) *gin.RouterGroup {
	group := router.Group("/api")
	group.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})

	group.GET("/todos", apiCtx.TodoAPI.FindAll)
	return group
}
