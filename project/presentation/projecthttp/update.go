package projecthttp

import (
	"clockify/project/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (projectcontroller *ProjectController) Update(c *gin.Context) {
	projectIDStr := c.Param("projectID")
	userIDStr := c.Param("userID")

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Project ID"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid User ID"})
		return
	}

	// Parse the JSON request body into update time entry request struct
	var updateProjectRequest domain.Projects
	if err := c.ShouldBindJSON(&updateProjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	// Validate required fields for the update
	if updateProjectRequest.Name == "" && updateProjectRequest.Client == "" && updateProjectRequest.Amount == 0 && updateProjectRequest.Tracked == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "at least one field is required for the update"})
		return
	}

	// Validate StartTime and EndTime if provided

	// Set the ID fields of updateTimeEntryRequest with the values from the path
	updateProjectRequest.ID = projectID
	updateProjectRequest.UserID = &userID

	// Call the service method to update the time entry
	if err := projectcontroller.projectService.Update(updateProjectRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update project"})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "project updated successfully"})
}
