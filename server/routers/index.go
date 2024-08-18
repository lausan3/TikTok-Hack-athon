package routers

import (
	"database/sql"
	routers "main/routers/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB, redisClient *redis.Client) {
	// 404 Route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	routers.RegisterAPIRoutes(router, db, redisClient)
}
