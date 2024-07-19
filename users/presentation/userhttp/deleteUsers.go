package userhttp

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (userController *UserController) Delete(c *gin.Context) {
	userId := c.Param("userID")
	if userId <= "" {
		fmt.Println("userId is Invalid")
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Attempt to delete the user
	err = userController.userService.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"error": "Failed to delete the user"})
		return
	} else if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something is missing! user is not deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User and associated projects and time-entries deleted successfully"})
}
