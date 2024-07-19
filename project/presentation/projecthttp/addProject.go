package projecthttp

import (
	"clockify/project/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
)

type ProjectController struct {
	projectService domain.ProjectService
}

func NewProjectController(projectService domain.ProjectService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}

func (projectcontroller *ProjectController) Create(ctx *gin.Context) {
	userIDstr := ctx.Param("userID")
	if userIDstr == "" {
		fmt.Errorf("userID is Invalid")
	}
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID is Invalid"})
		return
	}
	var createProjectRequest domain.Projects

	if err := ctx.ShouldBindJSON(&createProjectRequest); err != nil {
		var errorMsg string
		if verr, ok := err.(validator.ValidationErrors); ok {
			var fields []string
			for _, fieldErr := range verr {
				fieldName := fieldErr.StructField()
				fields = append(fields, fieldName)
			}
			errorMsg = fmt.Sprintf("Validation Errors: %s", strings.Join(fields, ", "))
		} else {
			errorMsg = "Failed to parse JSON"
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
		return
	}

	createProjectRequest.UserID = &userID
	fmt.Println(createProjectRequest)

	projectID, err := projectcontroller.projectService.Create(createProjectRequest)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Duplicate entry"):
			ctx.JSON(http.StatusConflict, gin.H{"error": "one field is missing"})

		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		fmt.Println("Error:", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": "project created successfully", "projectID": projectID.ID})

}
