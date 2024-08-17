package routers

import (
	"database/sql"

	controllers "main/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.RouterGroup, db *sql.DB) {
	router.POST("post", func(c *gin.Context) {
		controller := new(controllers.PostController)
		controller.CreatePost(c, db)
	})
}
