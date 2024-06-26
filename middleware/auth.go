package middleware

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware function that checks the Authorization header
// for a valid JWT token. If the token is valid, it sets the user ID in the context
// and allows the request to proceed. If the token is missing or invalid, it aborts
// the request with an unauthorized status.
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}

	context.Set("userId", userId)
	context.Next()
}
