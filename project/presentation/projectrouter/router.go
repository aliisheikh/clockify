package projectrouter

import (
	"clockify/project/presentation/projecthttp"
	"github.com/gin-gonic/gin"
)

func ProjectRouter(projectRouter projecthttp.ProjectController) *gin.Engine {

	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(200, "Welcome to the home page")
	})

	api := router.Group("/api")

	user := api.Group("/users")

	project := user.Group("/:userID/projects")

	{

		project.POST("", projectRouter.Create)
		project.DELETE(":projectID", projectRouter.Delete)
		project.GET(":projectID", projectRouter.GetProjectByID)
		project.GET("", projectRouter.GetByUserID)
		project.PUT(":projectID", projectRouter.Update)

	}
	return router
}
