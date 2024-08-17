package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRoutes(router *gin.Engine) {
	index := router.Group("/")

	index.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "LETS GOOOO",
		})
	})
}
