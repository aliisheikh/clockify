package timehttp

import (
	"clockify/timeEntry/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//
//type StartTimeEntryRequest struct {
//	Description string `json:"description"` // Add any other necessary fields
//}

type StartTimeEntryResponse struct {
	Message   string            `json:"message"`
	TimeEntry *domain.TimeEntry `json:"time_entry"`
}

func (timeEntryController *TimeEntryController) StartTimeEntry(c *gin.Context) {
	userIDStr := c.Param("userID")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID is Invalid"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID is Invalid"})
		return
	}

	var req domain.TimeEntry
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		fmt.Println("first error", err.Error())
		return
	}

	timeEntry, err := timeEntryController.timeEntryService.StartTimeEntry(uint(userID), req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		fmt.Println("second error", err.Error())
		return
	}

	c.JSON(http.StatusOK, StartTimeEntryResponse{
		Message:   "Time entry started successfully",
		TimeEntry: timeEntry,
	})
}
