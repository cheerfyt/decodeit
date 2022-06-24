package pkg

import (
	"embed"
	"net/http"
)

type Server struct {
	Address string
	Handler http.Handler
}

func NewServer(addr string, webFS embed.FS) *Server {
	h := NewHandler(webFS).Handle()
	return &Server{
		Address: addr,
		Handler: h,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.Address, s.Handler)
}
