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

func ApiFour(c *gin.Context, logger *zerolog.Logger) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API 4 Success. Authentication and Authorization required with level2 access.",
	})
}

func RegisterAPI4Route(router *gin.Engine, database *gorm.DB, logger *zerolog.Logger) {
	router.GET("/apifour",
		func(c *gin.Context) {
			middlewares.RequireAuthentication(c, database)
		},
		func(c *gin.Context) {
			middlewares.RequireAuthorization(c, "level2")
		},
		func(c *gin.Context) {
			ApiFour(c, logger)
		})
}
