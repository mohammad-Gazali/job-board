package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/mohammad-gazali/job-board/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	s := &Server{
		port: port,
		db: database.New(),
	}

	return &http.Server{
		Addr: fmt.Sprintf(":%d", s.port),
		Handler: s.RegisterRoutes(),
		// you can add IdleTimeout, ReadTimeout and WriteTimeout
	}
}