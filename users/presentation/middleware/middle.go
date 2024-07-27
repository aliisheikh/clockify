package middleware

import (
	"clockify/users/domain"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix from token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &domain.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("yoursecretkey"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		fmt.Printf("extrated userID from Token %d\n", claims.ID)
		// Ensure userID is set as an int
		c.Set("userID", claims.ID)
		c.Next()
	}
}
