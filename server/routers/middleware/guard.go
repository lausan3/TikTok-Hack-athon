package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("admin_key")
		secret := viper.GetString("API_ADMIN_KEY")

		if role != secret {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
