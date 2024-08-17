package routers

import (
	"database/sql"

	controllers "main/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.RouterGroup, db *sql.DB) {
	router.POST("posts/:username", func(c *gin.Context) {
		controller := new(controllers.PostController)
		controller.CreatePost(c, db)
	})

	router.GET("posts/:username", func(c *gin.Context) {
		controller := new(controllers.PostController)
		controller.GetPostsByUser(c, db)
	})
}
