package routes

import (
	"BizMart/pkg/controllers/categoryControllers"
	"BizMart/pkg/controllers/handlers"
	middlewares2 "BizMart/pkg/controllers/middlewares"
	"BizMart/pkg/controllers/orderControllers"
	"BizMart/pkg/controllers/productControllers"
	"BizMart/pkg/controllers/storesControllers"
	"BizMart/pkg/controllers/usersControllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	// usersRoute Маршруты для пользователей (профили)
	usersRoute := r.Group("/users")
	{
		usersRoute.GET("", usersControllers.GetAllUsers)
		usersRoute.GET("/:id", usersControllers.GetUserByID)
	}

	// auth Маршруты для авторизаций
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", usersControllers.SignUp)
		auth.POST("/sign-in", usersControllers.SignIn)
	}

	// storeRoutes Маршруты для магазинов
	storeRoutes := r.Group("/stores")
	{
		storeRoutes.GET("/", storesControllers.GetStores)
		storeRoutes.GET("/:id", storesControllers.GetStoreByID)
		storeRoutes.POST("/", middlewares2.CheckUserAuthentication, storesControllers.CreateStore)
		storeRoutes.PUT("/:id", middlewares2.CheckUserAuthentication, storesControllers.UpdateStore)
		storeRoutes.DELETE("/:id", middlewares2.CheckUserAuthentication, storesControllers.DeleteStore)
	}

	// reviewRoutes Маршруты для отзывов на магазины
	reviewRoutes := r.Group("/reviews")
	{
		reviewRoutes.GET("/")
		reviewRoutes.GET("/:id")
		reviewRoutes.POST("/", middlewares2.CheckUserAuthentication)
		reviewRoutes.PUT("/:id", middlewares2.CheckUserAuthentication)
		reviewRoutes.DELETE("/:id", middlewares2.CheckUserAuthentication)
	}

	r.GET("/hash-password", middlewares2.CheckSecretKey, handlers.HashPassword)

	// categoryRoutes Маршруты для категорий на магазине
	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.GET("/", categoryControllers.GetAllCategories)
		categoryRoutes.GET("/:id", categoryControllers.GetCategoryById)
		categoryRoutes.POST("/", middlewares2.CheckUserAuthentication, middlewares2.CheckAdmin, categoryControllers.CreateCategory)
		categoryRoutes.PUT("/:id", middlewares2.CheckUserAuthentication, middlewares2.CheckAdmin, categoryControllers.UpdateCategory)
		categoryRoutes.DELETE("/:id", middlewares2.CheckUserAuthentication, middlewares2.CheckAdmin, categoryControllers.DeleteCategory)
	}

	// orderStatusGroup Маршруты для статусов заказов
	orderStatusGroup := r.Group("/order-status")
	{
		orderStatusGroup.GET("/", orderControllers.GetAllOrderStatuses)
		orderStatusGroup.GET("/:id", orderControllers.GetOrderStatusByID)
		orderStatusGroup.POST("/", middlewares2.CheckUserAuthentication, middlewares2.CheckAdmin, orderControllers.CreateOrderStatus)
		orderStatusGroup.PUT("/:id", middlewares2.CheckUserAuthentication, middlewares2.CheckAdmin, orderControllers.UpdateOrderStatus)
		orderStatusGroup.DELETE("/:id", middlewares2.CheckUserAuthentication, middlewares2.CheckAdmin, orderControllers.DeleteOrderStatus)
	}

	// Обработчик статусов заказов по имени
	r.GET("/order-status/name/:name", orderControllers.GetOrderStatusByName)

	productGroup := r.Group("/product")
	{
		productGroup.GET("/", productControllers.GetAllProducts)
		productGroup.GET("/:id", productControllers.GetProductByID)
		productGroup.POST("/", middlewares2.CheckUserAuthentication, productControllers.CreateProduct)
	}
	return r
}
