package initializers

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func LoadEnvVariables(logger *zerolog.Logger) {
	err := godotenv.Load()
	if err != nil {
		logger.Error().Msg("Error loading .env file")
	}
	logger.Info().Msg("Environment variables loaded successfully.")
}
