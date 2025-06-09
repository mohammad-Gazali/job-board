package server

import (
	"fmt"
	"net/http"

	"github.com/mohammad-gazali/job-board/internal/repository"
)

type Server struct {
	config *ServerConfig
}

type ServerConfig struct {
	Port int
	Queries *repository.Queries
}

func NewServer(config *ServerConfig) *http.Server {
	s := &Server{ config: config }

	return &http.Server{
		Addr: fmt.Sprintf(":%d", s.config.Port),
		Handler: s.RegisterRoutes(),
		// you can add IdleTimeout, ReadTimeout and WriteTimeout
	}
}