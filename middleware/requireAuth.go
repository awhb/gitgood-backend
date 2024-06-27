package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/awhb/gitgood-backend/initialisers"
	"github.com/awhb/gitgood-backend/models"
)

func requireAuth(c *gin.Context) {
	// Get the cookie off request
	tokenString, err := c.Cookie("Authorisation")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiry
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the uer with token sub
		var user models.User
		initialisers.DB.First(&user, claims["subject"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to request
		c.Set("user", user)

		// Continue
		c.Next()
	}


}