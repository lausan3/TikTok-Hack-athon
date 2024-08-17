package routers

import (
	"database/sql"

	controllers "main/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, db *sql.DB) {
	router.POST("users", func(c *gin.Context) {
		controller := new(controllers.UserController)
		controller.RegisterUser(c, db)
	})

	router.DELETE("users/:username", func(c *gin.Context) {
		controller := new(controllers.UserController)
		controller.DeleteUser(c, db)
	})

	router.GET("users/:username", func(c *gin.Context) {
		controller := new(controllers.UserController)
		controller.GetUser(c, db)
	})
}
