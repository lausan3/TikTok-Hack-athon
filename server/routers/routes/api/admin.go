package routers

import (
	"database/sql"

	controllers "main/controllers"
	middleware "main/routers/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, db *sql.DB) {
	group := router.Group("/admin")
	group.Use(middleware.AdminMiddleware())

	group.DELETE("users/:username", func(c *gin.Context) {
		controller := new(controllers.UserController)
		controller.DeleteUser(c, db)
	})
}
