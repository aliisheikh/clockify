package userhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (userController *UserController) GetAllUsers(ctx *gin.Context) {
	// Call the service method to fetch all users
	users, err := userController.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users", "Internal": err.Error()})
		return
	}

	// If no users are found, return an appropriate message

	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}
	//resp := adapter.DomainToUser()

	// Respond with the list of users
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
