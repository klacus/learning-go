package routes

// This file contains the route definitions for a sample API that would require an authenticated and authorized user.
// This route is needs authentication, but not authorization

import (
	middleware "jwt/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func ApiTwo(c *gin.Context, logger *zerolog.Logger) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API 2. Authentication required.",
	})
}

func RegisterAPI2Route(router *gin.Engine, database *gorm.DB, logger *zerolog.Logger) {
	router.GET("/apitwo",
		func(c *gin.Context) {
			middleware.RequireAuthentication(c, database, logger)
		},
		func(c *gin.Context) {
			ApiTwo(c, logger)
		})
}
