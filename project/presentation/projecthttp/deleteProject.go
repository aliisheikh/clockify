package projecthttp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (projectcontroller *ProjectController) Delete(c *gin.Context) {
	projectIDStr := c.Param("projectID")
	userIDStr := c.Param("userID")

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Project ID"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid User ID"})
		return
	}

	err = projectcontroller.projectService.Delete(userID, projectID)
	if err != nil {
		if errors.Is(err, errors.New("project not found")) {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to delete project", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted project"})
}
