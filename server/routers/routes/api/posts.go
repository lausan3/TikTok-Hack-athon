package routers

import (
	"database/sql"

	controllers "main/controllers"
	"main/routers/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.RouterGroup, db *sql.DB) {
	postController := new(controllers.PostController)

	// get all posts
	router.GET("posts", func(c *gin.Context) {
		postController.GetAllPosts(c, db)
	})

	router.POST("posts/:username", middleware.JwtAuthMiddleware(), func(c *gin.Context) {
		postController.CreatePost(c, db)
	})

	router.GET("posts/:username", func(c *gin.Context) {
		postController.GetPostsByUser(c, db)
	})

}
