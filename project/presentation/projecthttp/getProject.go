package projecthttp

import (
	"clockify/project/presentation/adapter"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (projectcontroller *ProjectController) GetProjectByID(c *gin.Context) {
	projectID := c.Param("projectID")
	userID := c.Param("userID")

	projIDInt, err := strconv.Atoi(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Project ID"})
		return
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	project, err := projectcontroller.projectService.GetProjectByID(userIDInt, projIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Project"})
		return
	}

	if project == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
	}
	resp := adapter.DomainToProject(*project)

	c.JSON(http.StatusOK, gin.H{"data": resp})

}
