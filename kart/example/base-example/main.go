package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handler := gin.Default()
	app, cleanup, err := wireApp(handler)
	if err != nil {
		fmt.Println(111)
		panic(err)
	}
	defer cleanup()

	ctx := context.WithValue(context.Background(), "223", "122")
	if err = app.Run(ctx); err != nil {

	}
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-chSig
	fmt.Println("Server exiting")
}
