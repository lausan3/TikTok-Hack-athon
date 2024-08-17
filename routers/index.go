package routers

import (
	"database/sql"
	routers "main/routers/routes"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.LoadHTMLGlob("site/templates/*")

	// Load scripts
	router.Static("/site/scripts/", "./site/scripts/")

	// 404 Route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	routers.IndexRoutes(router)
	routers.DatabaseRoutes(router, db)
}
