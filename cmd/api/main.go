package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/mohammad-gazali/job-board/internal/server"
)

func main() {
	s := server.NewServer()

	slog.Info(fmt.Sprintf("Running on %s", s.Addr))

	log.Fatal(s.ListenAndServe())
}