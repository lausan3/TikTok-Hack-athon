package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.Engine) {
	index := router.Group("/posts")

	index.GET("/create-post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create-post.html", gin.H{})
	})
}
