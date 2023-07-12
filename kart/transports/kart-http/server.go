package kart_http

import (
	"context"
	"fmt"
	"github.com/costa92/errors"
	"net/http"
	"time"
)

type Server struct {
	Config      *HttpConfig
	Handler     http.Handler
	Middlewares []string
}

func NewServer(opts ...Option) *Server {
	srv := &Server{}
	for _, o := range opts {
		o(srv)
	}
	return srv
}

func (s *Server) Run(ctx context.Context) error {
	serverConfig := s.Config
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", serverConfig.Port),
		Handler:        s.Handler,
		ReadTimeout:    time.Duration(serverConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(serverConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) { // 如果是关闭状态，不当异常处理
			return err
		}
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
