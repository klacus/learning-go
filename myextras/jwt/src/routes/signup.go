package routes

import (
	"jwt/controllers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func RegisterSignupRoute(router *gin.Engine, database *gorm.DB, logger *zerolog.Logger) {
	router.POST("/signup", func(c *gin.Context) {
		controllers.Signup(c, database, logger)
	})
}
