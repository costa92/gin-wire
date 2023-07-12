package kart_http

import (
	"context"
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_NewServer(t *testing.T) {
	handler := gin.Default()
	config := &HttpConfig{}
	src := NewServer(
		WithHandler(handler),
		WithConfig(config),
	)
	if src == nil {
		t.Error("Server is nil")
	}
	ctx := context.WithValue(context.Background(), "test", "test")
	src.Run(ctx)
}
