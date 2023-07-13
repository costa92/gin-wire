package kart_grpc

import (
	"context"
	"testing"
	"time"
)

type testKey struct{}

func TestServer(t *testing.T) {
	config := &GrpcConfig{
		Port: "8081",
		Addr: "0.0.0.0",
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, testKey{}, "test")
	srv := NewGrpcServer(WithConfig(config))

	if e, err := srv.Endpoint(); err != nil || e == nil {
		t.Fatal(e, err)
	}

	go func() {
		// start server
		if err := srv.Start(ctx); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
}
