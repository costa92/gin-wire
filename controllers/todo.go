package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var TodoAPISet = wire.NewSet(ProvideTodoAPI)

type TodoAPI struct {
	MasterDB *gorm.DB
}

func ProvideTodoAPI() *TodoAPI {
	return &TodoAPI{}
}

func (t *TodoAPI) FindAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": "todo"})
}
