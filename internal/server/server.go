package server

import (
	"fmt"
	"net/http"

	"github.com/mohammad-gazali/job-board/internal/db"
)

type Server struct {
	config *ServerConfig
}

type ServerConfig struct {
	Port int
	Queries *db.Queries
}

func NewServer(config *ServerConfig) *http.Server {
	s := &Server{ config: config }

	return &http.Server{
		Addr: fmt.Sprintf(":%d", s.config.Port),
		Handler: s.RegisterRoutes(),
		// you can add IdleTimeout, ReadTimeout and WriteTimeout
	}
}