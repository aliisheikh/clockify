package userhttp

import (
	"clockify/users/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (userController *UserController) Register(ctx *gin.Context) {
	var req domain.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := userController.userService.Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error2": "username already exist"})
		return
	}

	ctx.JSON(http.StatusCreated, RegisterResponse{
		Message: "user signUp successfully",
		UserID:  uint(createdUser.ID),
	})
}
