package main

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var Database *gorm.DB // Global variable to hold the database connection

func ConnectToSQLite() (*gorm.DB, error) {
	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		dbName = "jwt.db"
	}

	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return database, nil
}

// type SQLLiteDB struct {
// 	database *gorm.DB
// 	logger   *zerolog.Logger
// }

// func ConnectToSQLite(logger *zerolog.Logger) (*SQLLiteDB, error) {
// 	dbName := os.Getenv("DBNAME")
// 	if dbName == "" {
// 		dbName = "jwt.db"
// 	}

// 	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
// 	if err != nil {
// 		logger.Error().Msgf("Error connecting to SQLite database: %s", err)
// 		return nil, err
// 	}
// 	logger.Info().Msgf("Connected to SQLite database: %s", dbName)

// 	return &SQLLiteDB{database: database, logger: logger}, nil
// }

// func (db *SQLLiteDB) Migrate() error {
// 	// Migrate the schema
// 	if err := db.database.AutoMigrate(&models.User{}); err != nil {
// 		db.logger.Error().Msgf("Error migrating database schema: %s", err)
// 	}
// 	db.logger.Info().Msg("Database migration completed successfully.")
// 	return nil
// }
