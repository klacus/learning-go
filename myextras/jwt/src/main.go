package main

import (
	"fmt"
	"jwt/initializers"
	"os"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
var database *gorm.DB // Global variable to hold the database connection

func init() {
	initializers.LoadEnvVariables(&logger)

	var err error
	database, err = ConnectToSQLite()
	if err != nil {
		logger.Error().Msgf("Error connecting to database for schema synchronization: %s", err)
		return
	}

	initializers.SyncDatabase(database, &logger)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiServer := NewApiServer(fmt.Sprintf(":%s", port))
	if err := apiServer.Start(); err != nil {
		logger.Error().Msgf("Error starting server: %s", err)
		return
	}
	apiServer.Start()
}
