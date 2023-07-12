package kart_http

import (
	"github.com/gin-gonic/gin"
)

type Option func(server *Server)

func WithConfig(config *HttpConfig) Option {
	return func(s *Server) {
		s.Config = config
	}
}

func WithGinEngin(engin *gin.Engine) Option {
	return func(s *Server) {
		s.GinEngin = engin
	}
}
