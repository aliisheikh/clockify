package userhttp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMeResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetMe godoc
// @Summary       Get current user information
// @Description   Retrieve information about the currently authenticated user
// @Produce json
// @Success 200 {object} GetMeResponse "Successful response"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /users/me [get]
func (userController *UserController) GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	fmt.Println(userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Debugging output to verify the user ID type
	fmt.Printf("Retrieved userID from context: %v, type: %T\n", userID, userID)

	// Assert userID is an int
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := userController.userService.GetUserByID(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response := GetMeResponse{
		ID:       uint(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, response)
}
