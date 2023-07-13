package kart_http

import (
	"context"
	"fmt"
	"github.com/costa92/errors"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"kart-cloud/kart/transports/kart-http/middlewares"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Config     *HttpConfig
	GinEngin   *gin.Engine
	httpServer *http.Server
}

func NewServer(opts ...Option) *Server {
	srv := &Server{}
	for _, o := range opts {
		o(srv)
	}
	srv.initAPIServer()
	return srv
}

func (s *Server) initAPIServer() {
	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

func (s *Server) Setup() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func (s *Server) InstallMiddlewares() {
	for _, m := range s.Config.Middlewares {
		mw, ok := middlewares.Middlewares[m]
		if !ok {
			continue
		}
		s.GinEngin.Use(mw)
	}
}

func (s *Server) InstallAPIs() {
	// Healthz 检测健康
	if s.Config.Healthz {
		s.GinEngin.GET("/healthz", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, http.StatusText(http.StatusOK))
		})
	}

	// install metric handler
	if s.Config.EnableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.GinEngin)
	}
}

func (s *Server) Start(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("appService recover err", "err", err)
		}
	}()
	serverConfig := s.Config
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", serverConfig.Port),
		Handler:        s.GinEngin,
		ReadTimeout:    time.Duration(serverConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(serverConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("start run http server:", serverConfig.Port)
	if err := s.httpServer.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) { // 如果是关闭状态，不当异常处理
			log.Print("start run failed server:", serverConfig.Port)
			return err
		}
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	log.Print("[HTTP] server stopping")
	return s.httpServer.Shutdown(ctx)
}
