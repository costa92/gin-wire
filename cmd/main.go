package main

import (
	"context"
	"costa92/gin-wire/config"
	"costa92/gin-wire/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionConfig().EncoderConfig),
			&zaptest.Discarder{},
			zap.DebugLevel,
		))
	conf := &config.Configuration{
		Server: config.ServerConfig{
			Port: "8080",
		},
		MasterDB: pkg.MySQLConfig{
			Dsn:          "",
			MaxOpenCount: 10,
		},
	}
	app, cleanup, err := wireApp(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 启动应用
	log.Print("start app ...")
	if err = app.Run(); err != nil {
		panic(err)
	}

	// 等待中断信号以优雅地关闭应用
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Print("shutdown app  ...")

	// 设置 5 秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭应用
	if err := app.Stop(ctx); err != nil {
		panic(err)
	}
}
