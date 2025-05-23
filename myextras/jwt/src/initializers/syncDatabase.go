package initializers

import (
	"jwt/models"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SyncDatabase(database *gorm.DB, logger *zerolog.Logger) {
	// Sync the database schema
	if err := database.AutoMigrate(&models.AccessLevel{}); err != nil {
		logger.Error().Msgf("Error migrating database schema for AccessLevel: %s", err)
		return
	}
	if err := database.AutoMigrate(&models.User{}); err != nil {
		logger.Error().Msgf("Error migrating database schema for User: %s", err)
		return
	}

	// Seed the database with initial example data
	// Yes sample passwords are there for example users. This is not a real application, just a learning example.

	// Create example access levels
	accessLevels := []models.AccessLevel{
		{Name: "level1", Description: "Level 1 access"},
		{Name: "level2", Description: "Level 2 access"},
		{Name: "level3", Description: "Level 3 access"},
	}
	result := database.FirstOrCreate(&accessLevels[0], accessLevels[0])
	result = database.FirstOrCreate(&accessLevels[1], accessLevels[1])
	result = database.FirstOrCreate(&accessLevels[2], accessLevels[2])
	if result.Error != nil {
		logger.Error().Msgf("Error seeding access levels: %s", result.Error)
	}

	// Create example users with hashed default passwords
	password1, _ := bcrypt.GenerateFromPassword([]byte("password1"), bcrypt.DefaultCost)
	password2, _ := bcrypt.GenerateFromPassword([]byte("password2"), bcrypt.DefaultCost)
	users := []models.User{
		{Name: "user1", Email: "email1@example.com", Password: string(password1)},
		{Name: "user2", Email: "email2@example.com", Password: string(password2)},
	}
	resultUsers := database.Where("email = ?", users[0].Email).FirstOrCreate(&users[0])
	if resultUsers.Error != nil {
		logger.Error().Msgf("Error seeding user1: %s", resultUsers.Error)
	}
	resultUsers = database.Where("email = ?", users[1].Email).FirstOrCreate(&users[1])
	if resultUsers.Error != nil {
		logger.Error().Msgf("Error seeding user2: %s", resultUsers.Error)
	}

	// Map users to access levels
	resultUserAccess := database.Model(&users[0]).Association("AccessLevels").Append(&accessLevels[0])
	if resultUserAccess != nil {
		logger.Error().Msgf("Error mapping user %s to access level %s: %s", users[0].Name, accessLevels[0].Name, resultUserAccess)
	}
	resultUserAccess = database.Model(&users[1]).Association("AccessLevels").Append(&accessLevels[1])
	if resultUserAccess != nil {
		logger.Error().Msgf("Error mapping user %s to access level %s: %s", users[1].Name, accessLevels[1].Name, resultUserAccess)
	}

	// Log the successful synchronization
	logger.Info().Msg("Database schema synchronized and seeded successfully.")
}
