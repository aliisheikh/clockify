package timehttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (timeEntryController *TimeEntryController) GetByUserID(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
		return
	}

	timeEntries, err := timeEntryController.timeEntryService.GetByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch time entries", "internalError": err.Error()})
		return
	}

	if len(timeEntries) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No time entries found for this user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": timeEntries})
}
