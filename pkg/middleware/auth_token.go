package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/EZChain-core/price-service/config"

	"github.com/EZChain-core/price-service/logger"


)


func AuthenticationTokenRequired(appConfig *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		xIAMToken := c.GetHeader("x-iam-token")
		if xIAMToken == "" || xIAMToken != appConfig.IAMToken {
			logger.Debug("Unauthorized, iam token is missing or invalid")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error_code": http.StatusUnauthorized,
				"message": "Unauthorized, Access token is missing or invalid",
				"success": false,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
