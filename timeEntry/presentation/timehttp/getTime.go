package timehttp

import (
	"clockify/timeEntry/presentation/adapter"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (timeEntryController *TimeEntryController) GetTimeEntryByID(ctx *gin.Context) {
	timeEntryIDStr := ctx.Param("timeEntryID")
	userIDStr := ctx.Param("userID")

	timeEntryID, err := strconv.Atoi(timeEntryIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Time Entry ID"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	timeEntry, err := timeEntryController.timeEntryService.GetTimeEntryByID(userID, timeEntryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Time Entry"})
		return
	}

	if timeEntry == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Time Entry not found"})
		return
	}

	// Assuming there's an adapter to convert domain.TimeEntry to a response format
	resp := adapter.DomainToTime(*timeEntry)

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
