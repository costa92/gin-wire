package routers

import "github.com/gin-gonic/gin"

func setApiGroupRoutes(router *gin.Engine) *gin.RouterGroup {
	group := router.Group("/api")
	group.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})
	return group
}
