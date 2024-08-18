package routers

import (
	"database/sql"

	controllers "main/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, db *sql.DB) {
	group := router.Group("/auth")
	authController := new(controllers.AuthController)

	group.POST("login", func(c *gin.Context) {
		authController.LoginUser(c, db)
	})

	router.POST("signup", func(c *gin.Context) {
		authController.RegisterUser(c, db)
	})
}
