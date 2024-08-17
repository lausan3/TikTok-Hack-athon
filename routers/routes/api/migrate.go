package routers

import (
	"database/sql"
	"main/db/migrations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MigrationRoutes(router *gin.RouterGroup, db *sql.DB) {
	database := router.Group("/migrate")

	database.GET("users", func(c *gin.Context) {
		migrations.MigrateUsersTable(db)

		c.JSON(http.StatusOK, gin.H{
			"message": "Users table migration was successful",
		})
	})
	database.GET("posts", func(c *gin.Context) {
		migrations.MigratePostsTable(db)

		c.JSON(http.StatusOK, gin.H{
			"message": "Posts table migration was successful",
		})
	})
}
