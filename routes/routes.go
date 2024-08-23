package routes

import (
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
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	// storeRoutes Маршруты для магазинов
	storeRoutes := r.Group("/stores")
	{
		storeRoutes.GET("/")
		storeRoutes.GET("/:id")
		storeRoutes.POST("/")
		storeRoutes.PUT("/:id")
		storeRoutes.DELETE("/:id")
	}

	// reviewRoutes Маршруты для отзывов на магазины
	reviewRoutes := r.Group("/reviews")
	{
		reviewRoutes.GET("/")
		reviewRoutes.GET("/:id")
		reviewRoutes.POST("/")
		reviewRoutes.PUT("/:id")
		reviewRoutes.DELETE("/:id")
	}
}
