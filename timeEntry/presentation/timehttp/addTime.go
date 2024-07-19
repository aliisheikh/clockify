package timehttp

import (
	domain2 "clockify/timeEntry/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
)

type TimeEntryController struct {
	timeEntryService domain2.TimeEntryService
}

func NewTimeEntryController(timeEntryService domain2.TimeEntryService) *TimeEntryController {
	return &TimeEntryController{
		timeEntryService: timeEntryService,
	}
}

func (timeEntryController *TimeEntryController) Create(ctx *gin.Context) {

	userIDstr := ctx.Param("userID")
	if userIDstr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID is Invalid"})
		return
	}

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID is Invalid"})
		return
	}

	var createTimeEntryRequest domain2.TimeEntry

	if err := ctx.ShouldBindJSON(&createTimeEntryRequest); err != nil {
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": errorMsg})
		return
	}

	createTimeEntryRequest.UserID = userID

	fmt.Println(createTimeEntryRequest)

	timeID, err := timeEntryController.timeEntryService.Create(createTimeEntryRequest)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Duplicate entry"):
			ctx.JSON(http.StatusConflict, gin.H{"error": "one field is missing"})

		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})

		}
		fmt.Println("Error:", err)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create Time Entry"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": "time entry created successfully", "timeID": timeID.ID})
}
