package projecthttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (projectcontroller *ProjectController) GetByUserID(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
		return
	}

	projects, err := projectcontroller.projectService.GetByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch projects", "Internal": err.Error()})
		return
	}
	if len(projects) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no project is assigned against this User"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": projects})
}
