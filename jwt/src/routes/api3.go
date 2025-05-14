package routes

import (
	"jwt/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// This file contains the route definitions for a sample API that would require an authenticated and authorized user.
// This route needs authentication and authorization to "level1" access.

func ApiThree(c *gin.Context, logger *zerolog.Logger) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API 3 Success. Authentication and Authorization required with level1 access.",
	})
}

func RegisterAPI3Route(router *gin.Engine, database *gorm.DB, logger *zerolog.Logger) {
	router.GET("/apithree",
		func(c *gin.Context) {
			middlewares.RequireAuthentication(c, database, logger)
		},
		func(c *gin.Context) {
			middlewares.RequireAuthorization(c, "level1", logger)
		},
		func(c *gin.Context) {
			ApiThree(c, logger)
		})
}
