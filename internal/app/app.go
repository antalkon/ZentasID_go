package app

import (
	"log/slog"

	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler"
	"github.com/antalkon/ZentasID_go/internal/transport/rest/router"
	"github.com/antalkon/ZentasID_go/pkg/config"
	"github.com/antalkon/ZentasID_go/pkg/connectDB"
	"github.com/antalkon/ZentasID_go/pkg/connectRedis"
	"github.com/antalkon/ZentasID_go/pkg/logger"
	z_validator "github.com/antalkon/ZentasID_go/pkg/validator"
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
		return // Exit if DB initialization fails
	}
	_ = connection

	// Initialize Redis
	connectRedis.InitRedis()

	// Initializate Validator
	z_validator.InitValidator()

	// Setup GIN
	h := handler.NewHandler()
	r := router.SetupRouter(h)

	if address == "" {
		log.Error("Failed start server: Address is empty.")
		return // Exit if address is not specified
	}

	if err := r.Run(address); err != nil {
		log.Error("Failed to run server. LOG:", err)
	}
}
