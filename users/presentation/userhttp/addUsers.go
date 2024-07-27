package userhttp

import (
	"clockify/users/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type RegisterResponse struct {
	Message string `json:"message"`
	UserID  uint   `json:"userID"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type UserController struct {
	userService domain.UserService
}

func NewUserController(userService domain.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// UserController.Create

// Swagger Data
// @Summary Create a new user
// @Description Register a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body domain.User true "User"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 409 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /users [post]
func (userController *UserController) Create(ctx *gin.Context) {
	// Initialize a User instance
	var createUserRequest domain.User
	fmt.Println("12121", createUserRequest)

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
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

	user, err := userController.userService.Create(createUserRequest)
	if err != nil {
		// Handle specific errors
		switch {
		case strings.Contains(err.Error(), "Duplicate entry"):
			ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	// Respond with a success message and include userID
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "userID": user.ID})
}
