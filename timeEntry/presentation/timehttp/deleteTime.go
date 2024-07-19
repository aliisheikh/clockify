package timehttp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (timeEntryController *TimeEntryController) Delete(c *gin.Context) {
	timeEntryIDStr := c.Param("timeEntryID")
	userIDStr := c.Param("userID")

	timeEntryID, err := strconv.Atoi(timeEntryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Time Entry ID"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid User ID"})
		return
	}

	err = timeEntryController.timeEntryService.Delete(userID, timeEntryID)
	if err != nil {
		if errors.Is(err, errors.New("time entry not found")) {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Time entry not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to delete time entry", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted time entry"})
}
