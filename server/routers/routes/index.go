package routers

import (
	"database/sql"
	endpoints "main/routers/routes/api"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine, db *sql.DB) {
	api := router.Group("/api")

	endpoints.MigrationRoutes(api, db)
	endpoints.AdminRoutes(api, db)
	endpoints.AuthRoutes(api, db)
	endpoints.UserRoutes(api, db)
	endpoints.PostRoutes(api, db)
}
