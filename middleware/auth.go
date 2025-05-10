package middleware

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strings"

	"mini-social-network/constants"
	"mini-social-network/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(constants.Authorization)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			c.Abort()
			return
		}

		tokenString := ""
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			tokenString = parts[1]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			c.Abort()
			return
		}

		userID, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
