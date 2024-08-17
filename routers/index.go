package routers

import (
	"database/sql"
	"main/db/migrations"
	routers "main/routers/routes"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.LoadHTMLGlob("site/templates/*")

	// 404 Route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })
	router.GET("/db/migrate/users", func(c *gin.Context) {
		migrations.MigrateUsersTable(db)

		c.JSON(http.StatusOK, gin.H{
			"message": "Users table migration was successful",
		})
	})
	router.GET("/db/migrate/posts", func(c *gin.Context) {
		migrations.MigratePostsTable(db)

		c.JSON(http.StatusOK, gin.H{
			"message": "Posts table migration was successful",
		})
	})

	routers.IndexRoutes(router)
}
