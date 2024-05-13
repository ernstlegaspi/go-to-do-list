package api

import (
	"net/http"
)

type server struct {
	addr string
}

func InitServer(addr string) *server {
	return &server{
		addr: addr,
	}
}

func (s *server) RunAPI() error {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return server.ListenAndServe()
}
