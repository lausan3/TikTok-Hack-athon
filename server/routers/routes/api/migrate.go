package routers

import (
	"database/sql"
	"main/db/migrations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MigrationRoutes(router *gin.RouterGroup, db *sql.DB) {
	group := router.Group("/migrate")

	group.GET("users", func(c *gin.Context) {
		migrations.MigrateUsersTable(db)

		c.JSON(http.StatusOK, gin.H{
			"message": "Users table migration was successful",
		})
	})
	group.GET("posts", func(c *gin.Context) {
		migrations.MigratePostsTable(db)

		c.JSON(http.StatusOK, gin.H{
			"message": "Posts table migration was successful",
		})
	})
}
