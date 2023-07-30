package Middleware

import (
	auth "datingapp/jwt"
	"github.com/gin-gonic/gin"
	"strings"
	_ "time"
)

func Auth() gin.HandlerFunc{
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, BEARER_SCHEMA)
		if authHeader == "" || tokenString == authHeader {
			context.JSON(401, gin.H{"responseCode":401, "responseMsg": "Not Authorized header type on Acces Token"})
			context.Abort()
			return
		}
		err:= auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(301, gin.H{"responseCode": 301, "responseMessage": "Access Token Invalid"})
			context.Abort()
			return
		}
		context.Next()
	}
}
