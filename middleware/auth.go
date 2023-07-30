package Middleware

import (
	auth "connector-permata/Auth"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/timeout"
	"strings"
	"time"
)

func Auth() gin.HandlerFunc{
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		// tokenString := authHeader[len(BEARER_SCHEMA):]
		tokenString := strings.TrimPrefix(authHeader, BEARER_SCHEMA)
		if authHeader == "" || tokenString == authHeader {
			context.JSON(401, gin.H{"error": "Not Authorized header type on Acces Token"})
			context.Abort()
			return
		}
		err:= auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"responseCode": "4017301", "responseMessage": "Access Token Invalid"})
			context.Abort()
			return
		}
		context.Next()
	}
}
