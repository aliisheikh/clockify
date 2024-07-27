package router

import (
	"clockify/users/presentation/middleware"
	"clockify/users/presentation/userhttp"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userRouter userhttp.UserController) *gin.Engine {
	router := gin.Default()

	//Add Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.POST("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.POST("/docs/*any, ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.POST("/docs/*any, ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.GET("/swagger/*any, ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.DELETE("/swagger/*any, ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.GET("/swagger/*any, ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		//user.GET("/me", userRouter.GetMe)
	}
	// authorized router
	authorized := api.Group("/users").Use(middleware.Authenticate())
	{
		authorized.GET("/me", userRouter.GetMe)
	}

	return router
}
