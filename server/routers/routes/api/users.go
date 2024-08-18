package routers

import (
	"database/sql"

	controllers "main/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, db *sql.DB) {
	userController := new(controllers.UserController)

	router.GET("users/:username", func(c *gin.Context) {
		userController.GetUser(c, db)
	})
}
