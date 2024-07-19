package userhttp

import (
	"clockify/users/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (userController *UserController) Login(ctx *gin.Context) {
	var req domain.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	match, err := userController.userService.CheckPassword(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username or password"})
		return
	}

	if match {
		token, err := userController.userService.GenerateToken(req.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		ctx.JSON(http.StatusOK, LoginResponse{Message: "login successfully", Token: token})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
