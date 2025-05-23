package routes

import (
	"jwt/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterLogoutRoute(router *gin.Engine) {
	router.DELETE("/logout", func(c *gin.Context) {
		controllers.Logout(c)
	})
}
