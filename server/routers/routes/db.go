package routers

import (
	"database/sql"
	"main/db/migrations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DatabaseRoutes(router *gin.Engine, db *sql.DB) {
	database := router.Group("/db")

	database.GET("migrate/users", func(c *gin.Context) {
		migrations.MigrateUsersTable(db)

		c.JSON(http.StatusOK, gin.H{
			"message": "Users table migration was successful",
		})
	})
	database.GET("migrate/posts", func(c *gin.Context) {
		migrations.MigratePostsTable(db)

		c.JSON(http.StatusOK, gin.H{
			"message": "Posts table migration was successful",
		})
	})
}
