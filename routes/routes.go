package routes

import (
	"BizMart/middlewares"
	"BizMart/pkg/controllers/Category"
	"BizMart/pkg/controllers/Order"
	"BizMart/pkg/controllers/Users"
	"BizMart/pkg/controllers/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// usersRoute Маршруты для пользователей (авторизация, профили)
	usersRoute := r.Group("/users")
	{
		usersRoute.GET("", Users.GetAllUsers)
		usersRoute.POST("", Users.CreateUser)
		usersRoute.GET("/:id", Users.GetUserByID)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", Users.SignUp)
		auth.POST("/sign-in", Users.SignIn)
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

	r.GET("/hash-password", middlewares.CheckSecretKey, handlers.HashPassword)

	// categoryRoutes Маршруты для категорий на магазине
	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.GET("/", Category.GetAllCategories)
		categoryRoutes.GET("/:id", Category.GetCategoryById)
		categoryRoutes.POST("/", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, Category.CreateCategory)
		categoryRoutes.PUT("/:id", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, Category.UpdateCategory)
		categoryRoutes.DELETE("/:id", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, Category.DeleteCategory)
	}

	// orderStatusGroup Маршруты для статусов заказов
	orderStatusGroup := r.Group("/orderstatus")
	{
		orderStatusGroup.GET("/", Order.GetAllOrderStatusses)
		orderStatusGroup.GET("/:id", Order.GetOrderStatusByID)
		orderStatusGroup.POST("/", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, Order.CreateOrderStatus)
		orderStatusGroup.PUT("/:id", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, Order.UpdateOrderStatus)
		orderStatusGroup.DELETE("/:id", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, Order.DeleteOrderStatus)
	}

	// Обработчик статусов заказов по имени
	r.GET("/orderstatus/name/:name", Order.GetOrderStatusByName)
}
