package routers

import (
	"context"
	"database/sql"
	"main/routers/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func Routes(ctx context.Context, db *sql.DB, redis *redis.Client) *gin.Engine {
	environment := viper.GetBool("DEBUG")

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	RegisterWebSocketRoutes(ctx, upgrader, router, redis)
	RegisterRoutes(router, db, redis)

	return router
}
