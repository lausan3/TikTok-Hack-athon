package routers

import (
	"database/sql"

	controllers "main/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, db *sql.DB) {
	router.POST("user", func(c *gin.Context) {
		controller := new(controllers.UserController)
		controller.RegisterUser(c, db)
	})

}
