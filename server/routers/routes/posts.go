package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.Engine) {
	posts := router.Group("/posts")

	posts.GET("/create-post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"content": "create-post.html",
		})
	})
}
