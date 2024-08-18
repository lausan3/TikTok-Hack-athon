package routers

import (
	"database/sql"
	endpoints "main/routers/routes/api"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RegisterAPIRoutes(router *gin.Engine, db *sql.DB, redisClient *redis.Client) {
	api := router.Group("/api")

	endpoints.MigrationRoutes(api, db)
	endpoints.AdminRoutes(api, db)
	endpoints.AuthRoutes(api, db)
	endpoints.UserRoutes(api, db)
	endpoints.PostRoutes(api, db, redisClient)
}
