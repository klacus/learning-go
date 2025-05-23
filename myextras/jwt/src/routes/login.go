package routes

import (
	"jwt/controllers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func RegisterLoginRoute(router *gin.Engine, database *gorm.DB, logger *zerolog.Logger) {
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, database, logger)
	})
}
