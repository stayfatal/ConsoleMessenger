package middleware

import (
	"messenger/internal/authentication"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		claims, err := authentication.ValidateToken(token)

		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		c.Set("id", claims.Id)

		c.Next()
	}
}
