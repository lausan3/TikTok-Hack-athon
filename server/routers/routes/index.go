package routers

import (
	"database/sql"
	endpoints "main/routers/routes/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRoutes(router *gin.Engine) {
	index := router.Group("/")

	index.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"content": "home.html",
		})
	})
}

func RegisterAPIRoutes(router *gin.Engine, db *sql.DB) {
	api := router.Group("/api")

	endpoints.MigrationRoutes(api, db)
	endpoints.UserRoutes(api, db)
	endpoints.PostRoutes(api, db)
}
