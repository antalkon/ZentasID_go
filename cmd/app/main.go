package main

import (
	"fmt"
	"log/slog"
	"os"

	log2 "log"

	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler"
	"github.com/antalkon/ZentasID_go/internal/transport/rest/router"

	"github.com/antalkon/ZentasID_go/pkg/config"
	_ "github.com/gin-gonic/gin"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	address := cfg.HTTPServer.Address
	log := setupLogger(cfg.Env)
	fmt.Println(cfg)
	log.Info("Starting app", slog.String("env", cfg.Env))

	h := handler.NewHandler()
	r := router.SetupRouter(h)

	if address == "" {
		log2.Fatal("Failed to run server: Adress is empty. \n pls. check config files")
	}

	if err := r.Run(address); err != nil {
		log2.Fatal("Failed to run server: %s", err)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
