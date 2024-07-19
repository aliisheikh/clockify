package userhttp

import (
	"clockify/users/presentation/adapter"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (userController *UserController) GetUserByID(c *gin.Context) {

	userId := c.Param("userID")
	fmt.Println(userId)
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	u, err := userController.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find the user"})
		return
	}

	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	resp := adapter.DomainToUser(*u)

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
