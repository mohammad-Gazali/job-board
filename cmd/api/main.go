package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/mohammad-gazali/job-board/internal/repository"
	"github.com/mohammad-gazali/job-board/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

// @title Documentation for Job Board
// @description API for dealing with job opportunities for both companies and employees sides
// @version 1

// @BasePath /api/v1
func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := repository.New(conn)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	config := &server.ServerConfig{
		Port: port,
		Queries: queries,
	}
	s := server.NewServer(config)

	slog.Info(fmt.Sprintf("Running on %s", s.Addr))

	return s.ListenAndServe()
}

func main() {
	log.Fatal(run())
}