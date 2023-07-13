package kart_grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"kart-cloud/kart/internal/host"
	"net"
	"net/url"
	"sync"
)

const HEALTHCHECK_SERVICE = "grpc.health.v1.Health"

type GrpcServer struct {
	config *GrpcConfig
	server *grpc.Server

	health   *health.Server
	listener net.Listener
	once     sync.Once
	err      error

	endpoint *url.URL
}

func NewGrpcServer(opts ...Option) *GrpcServer {
	g := &GrpcServer{
		health: health.NewServer(),
	}
	for _, opt := range opts {
		opt(g)
	}
	g.initGrpcServer()
	return g
}

func (s *GrpcServer) initGrpcServer() {
	var unaryInterceptors []grpc.UnaryServerInterceptor
	var streamInterceptors []grpc.StreamServerInterceptor

	grpcOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
		grpc.ChainStreamInterceptor(streamInterceptors...),
	}

	gs := grpc.NewServer(grpcOpts...)
	s.server = gs
	s.health.SetServingStatus(HEALTHCHECK_SERVICE, grpc_health_v1.HealthCheckResponse_SERVING)

	grpc_health_v1.RegisterHealthServer(s.server, s.health)
	reflection.Register(s.server)
}

// Endpoint return a real address to registry endpoint.
// examples:
// grpc://127.0.0.1:9000?isSecure=false
func (s *GrpcServer) Endpoint() (*url.URL, error) {
	config := s.config
	address := fmt.Sprintf("%s:%s", config.Addr, config.Port)
	s.once.Do(func() {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			s.err = err
			return
		}
		addr, err := host.Extract(address, s.listener)
		if err != nil {
			lis.Close()
			s.err = err
			return
		}
		s.listener = lis
		s.endpoint = &url.URL{Scheme: "grpc", Host: addr}
	})
	if s.err != nil {
		return nil, s.err
	}
	return s.endpoint, nil
}

func (s *GrpcServer) Start(ctx context.Context) error {
	if _, err := s.Endpoint(); err != nil {
		return err
	}
	fmt.Println("[gRPC] server started", "listen_addr", s.listener.Addr().String())
	s.health.Resume()
	return s.server.Serve(s.listener)
}

func (s *GrpcServer) Stop(ctx context.Context) error {
	return nil
}
