package main

import (
	"social-sys/internal/api"
	"social-sys/internal/config"
	"social-sys/internal/database"
	"social-sys/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.Load()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configurations")
	}

	db, err := database.Initialize(cfg.DatabaseURL)

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize database: %s", cfg.DatabaseURL)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to run database migrations")
	}

	router := gin.Default()

	api.SetupRoutes(router, db)

	err = router.Run(cfg.ServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
