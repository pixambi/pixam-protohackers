package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pixambi/pixam-protohackers/internal/config"
	"github.com/pixambi/pixam-protohackers/internal/server"
)

func main() {
	cfg := config.Load()

	slog.SetDefault(cfg.Logger)

	srv := server.New(cfg)
	if err := srv.Start(); err != nil {
		fmt.Println("server error: ", err.Error())
		os.Exit(1)
	}
}
