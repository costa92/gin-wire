package main

// https://github.com/jassue/gin-wire/blob/main/cmd/app/app.go
import (
	"context"
	"costa92/gin-wire/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	conf    *config.Configuration
	logger  *zap.Logger
	httpSrv *http.Server
}

func newApp(conf *config.Configuration, logger *zap.Logger, httpSrv *http.Server) *App {
	return &App{
		conf:    conf,
		logger:  logger,
		httpSrv: httpSrv,
	}
}

func (a *App) Run() error {
	a.logger.Info("Starting server")
	go func() {
		a.logger.Info("Server started")
		err := a.httpSrv.ListenAndServe()
		if err != nil {
			a.logger.Error("Server failed", zap.Error(err))
			panic(err)
		}
	}()
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("Stopping server")
	if err := a.httpSrv.Shutdown(ctx); err != nil {
		return err
	}
	a.logger.Info("Server stopped")
	return nil
}

func newHttpServer(conf *config.Configuration, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: router,
	}
}
