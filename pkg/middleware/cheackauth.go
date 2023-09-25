package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tegehhat/helper/pkg/handlers"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Need autorization"})
			c.Abort()
			return
		}

		session, err := handlers.GetSession(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		if time.Now().After(session.ExpiresAt) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Need reautorization"})
			c.Abort()
			return
		}

		// Передаем информацию о сессии в контекст
		c.Set("session", session)
		c.Next()
	}
}
