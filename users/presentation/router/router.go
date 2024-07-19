package router

import (
	"clockify/users/presentation/userhttp"
	"github.com/gin-gonic/gin"
)

func NewRouter(userRouter userhttp.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(200, "Welcome to the home page")
	})

	api := router.Group("/api")

	user := api.Group("/users")
	{
		user.POST("/register", userRouter.Register)
		user.POST("/login", userRouter.Login)
		user.POST("", userRouter.Create)
		user.DELETE(":userID", userRouter.Delete)
		user.GET(":userID", userRouter.GetUserByID)
		user.GET("", userRouter.GetAllUsers)
	}
	return router
}
