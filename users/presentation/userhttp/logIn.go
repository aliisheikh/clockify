package userhttp

import (
	"clockify/users/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Login a user
// @Description Authenticate a user and generate a token
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body domain.User true "User credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 401 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /users/login [post]
func (userController *UserController) Login(ctx *gin.Context) {
	var req domain.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		fmt.Println(err.Error())
		return
	}

	user, err := userController.userService.CheckPassword(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username or password"})
		fmt.Println(err.Error())
		return
	}

	if user != nil {
		token, err := userController.userService.GenerateToken(user.ID, user.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		ctx.JSON(http.StatusOK, LoginResponse{Message: "login successfully", Token: token})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
