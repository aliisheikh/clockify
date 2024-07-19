package timehttp

import (
	"clockify/timeEntry/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (timeEntryController *TimeEntryController) Update(c *gin.Context) {
	timeEntryIDStr := c.Param("timeEntryID")
	userIDStr := c.Param("userID")

	timeEntryID, err := strconv.Atoi(timeEntryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Time Entry ID"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	// Parse the JSON request body into update time entry request struct
	var updateTimeEntryRequest domain.TimeEntry
	if err := c.ShouldBindJSON(&updateTimeEntryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Validate required fields for the update
	if updateTimeEntryRequest.Description == "" && updateTimeEntryRequest.StartTime.IsZero() && updateTimeEntryRequest.EndTime.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one field is required for the update"})
		return
	}

	// Validate StartTime and EndTime if provided
	if !updateTimeEntryRequest.StartTime.IsZero() && !updateTimeEntryRequest.EndTime.IsZero() {
		if updateTimeEntryRequest.StartTime.After(updateTimeEntryRequest.EndTime) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "StartTime must be before EndTime"})
			return
		}
	}

	// Set the ID fields of updateTimeEntryRequest with the values from the path
	updateTimeEntryRequest.ID = timeEntryID
	updateTimeEntryRequest.UserID = userID

	// Call the service method to update the time entry
	if err := timeEntryController.timeEntryService.Update(updateTimeEntryRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update time entry"})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Time entry updated successfully"})
}
