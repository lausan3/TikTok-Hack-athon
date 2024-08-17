package routers

import (
	"database/sql"
	"main/routers/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routes(db *sql.DB) *gin.Engine {
	environment := viper.GetBool("DEBUG")

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

	RegisterRoutes(router, db)

	return router
}
