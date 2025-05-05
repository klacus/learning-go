package routes

import (
	middleware "jwt/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// This file contains the route definitions for a sample API that would require an authenticated and authorized user.
// This route needs authentication and authorization to "level1" access.

func ApiThree(c *gin.Context, logger *zerolog.Logger) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API 2. Authentication required.",
	})
}

func RegisterAPI3Route(router *gin.Engine, database *gorm.DB, logger *zerolog.Logger) {
	router.GET("/apithree",
		func(c *gin.Context) {
			middleware.RequireAuthentication(c, database)
		},
		func(c *gin.Context) {
			middleware.RequireAuthentication(c, database)
		},
		func(c *gin.Context) {
			ApiThree(c, logger)
		})
}
