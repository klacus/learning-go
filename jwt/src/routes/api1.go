package routes

// This file contains the route definitions for a sample API that would require an authenticated and authorized user.
// This route is public and does not need authentication or authorization
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ApiOne(c *gin.Context, logger *zerolog.Logger) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API 1. No authentication or authorization required.",
	})
}

func RegisterAPI1Route(router *gin.Engine, logger *zerolog.Logger) {
	router.GET("/apione", func(c *gin.Context) {
		ApiOne(c, logger)
	})
}
