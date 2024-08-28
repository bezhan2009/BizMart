package routes

import (
	"BizMart/middlewares"
	"BizMart/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// usersRoute Маршруты для пользователей (авторизация, профили)
	usersRoute := r.Group("/users")
	{
		usersRoute.GET("", controllers.GetAllUsers)
		usersRoute.POST("", controllers.CreateUser)
		usersRoute.GET(":id", controllers.GetUserByID)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
	}

	// storeRoutes Маршруты для магазинов
	storeRoutes := r.Group("/stores")
	{
		storeRoutes.GET("/")
		storeRoutes.GET("/:id")
		storeRoutes.POST("/", middlewares.CheckUserAuthentication)
		storeRoutes.PUT("/:id", middlewares.CheckUserAuthentication)
		storeRoutes.DELETE("/:id", middlewares.CheckUserAuthentication)
	}

	// reviewRoutes Маршруты для отзывов на магазины
	reviewRoutes := r.Group("/reviews")
	{
		reviewRoutes.GET("/")
		reviewRoutes.GET("/:id")
		reviewRoutes.POST("/", middlewares.CheckUserAuthentication)
		reviewRoutes.PUT("/:id", middlewares.CheckUserAuthentication)
		reviewRoutes.DELETE("/:id", middlewares.CheckUserAuthentication)
	}
}
