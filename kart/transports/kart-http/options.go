package kart_http

import "net/http"

type Option func(server *Server)

func WithConfig(config *HttpConfig) Option {
	return func(s *Server) {
		s.Config = config
	}
}

func WithHandler(handler http.Handler) Option {
	return func(s *Server) {
		s.Handler = handler
	}
}
