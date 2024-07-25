// internal/app/app.go
package app

import (
	"log/slog"

	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler"
	"github.com/antalkon/ZentasID_go/internal/transport/rest/router"
	"github.com/antalkon/ZentasID_go/pkg/config"
	"github.com/antalkon/ZentasID_go/pkg/connectDB"
	"github.com/antalkon/ZentasID_go/pkg/logger"
)

func Main() {
	cfg := config.MustLoad() // get server configuration
	address := cfg.HTTPServer.Address

	// Setup logger
	log := logger.SetupLogger(cfg.Env)
	log.Info("Logger session started.", slog.String("env", cfg.Env))
	log.Info("--- Server Started ---")

	// Initialize database
	connectDB.InitDB()
	connection, err := Db(cfg.Env, "development")
	if err != nil {
		log.Error("DB fatal error. Db not started. LOG:", err)
	}
	_ = connection
	// Setup GIN
	h := handler.NewHandler()
	r := router.SetupRouter(h)

	if address == "" {
		log.Error("Failed start server: Adres is no empty.")
	}

	if err := r.Run(address); err != nil {
		log.Error("Failed run server. LOG:", err)
	}
}
