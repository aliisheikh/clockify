package userhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Produce json
// @Success 200 {object} gin.H "List of users"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Failure 404 {object} gin.H "No users found"
// @Router /api/users [get]
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
