package middlewares

import (
	"log"
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		log.Println("not authorized: empty token")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		log.Printf("not authorized token: %v", err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized token"})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
