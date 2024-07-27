package userhttp

import (
	"clockify/users/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// Add Swagger
// @Summary Register a new user
// @Description Register a new user with a username, email, and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body domain.User true "User information"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 409 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /users/register [post]
func (userController *UserController) Register(ctx *gin.Context) {
	var req domain.User
	//if err := ctx.ShouldBindJSON(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
	//	return
	//}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var errorMsg string
		if verr, ok := err.(validator.ValidationErrors); ok {
			var fields []string
			for _, fieldErr := range verr {
				fieldName := fieldErr.StructField()
				fields = append(fields, fieldName)
			}
			errorMsg = fmt.Sprintf("Missing or invalid fields: %s", strings.Join(fields, ", "))
		} else {
			errorMsg = "Failed to parse JSON"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
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
