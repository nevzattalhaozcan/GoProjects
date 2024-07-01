package routes

import (
	"log"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		log.Printf("Error saving user: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user. try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		log.Printf("could not validate user: %v", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		log.Printf("something went wrong with token: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user logged in successfuly", "token": token})
}
