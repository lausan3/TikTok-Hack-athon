package routers

import (
	"database/sql"

	controllers "main/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.RouterGroup, db *sql.DB) {

	// get all posts
	router.GET("posts", func(c *gin.Context) {
		controller := new(controllers.PostController)
		controller.GetAllPosts(c, db)
	})

	router.POST("posts/:username", func(c *gin.Context) {
		controller := new(controllers.PostController)
		controller.CreatePost(c, db)
	})

	router.GET("posts/:username", func(c *gin.Context) {
		controller := new(controllers.PostController)
		controller.GetPostsByUser(c, db)
	})

}
