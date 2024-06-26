package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

// signup handles the user signup process.
// It binds the JSON payload to a User model, validates the input,
// and attempts to save the user to the database. Returns appropriate
// HTTP status codes and messages based on the result.
func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// login handles the user login process.
// It binds the JSON payload to a User model, validates the user's credentials,
// and generates a JWT token if the credentials are valid. Returns appropriate
// HTTP status codes and messages based on the result.
func login(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data"})
		return
	}

	err = user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Unauthorized"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successful login", "token": token})

}
