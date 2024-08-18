package routers

import (
	"database/sql"

	controllers "main/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, db *sql.DB) {
	group := router.Group("/auth")

	group.POST("login", func(c *gin.Context) {
		controller := new(controllers.AuthController)
		controller.LoginUser(c, db)
	})
}
