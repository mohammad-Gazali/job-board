package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/mohammad-gazali/job-board/internal/db"
	"github.com/mohammad-gazali/job-board/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

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