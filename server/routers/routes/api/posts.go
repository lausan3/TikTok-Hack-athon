package routers

import (
	"database/sql"

	controllers "main/controllers"
	"main/routers/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func PostRoutes(router *gin.RouterGroup, db *sql.DB, redisClient *redis.Client) {
	postController := new(controllers.PostController)

	// get all posts
	router.GET("posts", func(c *gin.Context) {
		postController.GetAllPosts(c, db)
	})

	router.POST("posts", middleware.JwtAuthMiddleware(), func(c *gin.Context) {
		postController.CreatePost(c, db, redisClient)
	})

	router.GET("posts/:username", func(c *gin.Context) {
		postController.GetPostsByUser(c, db)
	})

}
