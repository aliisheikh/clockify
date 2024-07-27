package timerouter

import (
	"clockify/timeEntry/presentation/timehttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TimeEntryRouter(timeEntryRouter timehttp.TimeEntryController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to the home page")
	})

	api := router.Group("/api")

	user := api.Group("/users")

	timeEntry := user.Group("/:userID/time-entries")

	{
		timeEntry.POST("", timeEntryRouter.Create)
		timeEntry.POST("/start", timeEntryRouter.StartTimeEntry)
		timeEntry.DELETE(":timeEntryID", timeEntryRouter.Delete)
		timeEntry.PUT(":timeEntryID", timeEntryRouter.Update)
		timeEntry.GET("", timeEntryRouter.GetByUserID)
		timeEntry.GET(":timeEntryID", timeEntryRouter.GetTimeEntryByID)
	}

	return router
}
