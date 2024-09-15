package routes

import (
	_ "BizMart/docs"
	"BizMart/internal/controllers"
	"BizMart/internal/controllers/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// usersRoute Маршруты для пользователей (профили)
	usersRoute := r.Group("/users")
	{
		usersRoute.GET("", controllers.GetAllUsers)
		usersRoute.GET("/:id", controllers.GetUserByID)
	}

	// auth Маршруты для авторизаций
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
	}

	// storeRoutes Маршруты для магазинов
	storeRoutes := r.Group("/store")
	{
		storeRoutes.GET("/", controllers.GetStores)
		storeRoutes.GET("/:id", controllers.GetStoreByID)
		storeRoutes.POST("/", middlewares.CheckUserAuthentication, controllers.CreateStore)
		storeRoutes.PUT("/:id", middlewares.CheckUserAuthentication, controllers.UpdateStore)
		storeRoutes.DELETE("/:id", middlewares.CheckUserAuthentication, controllers.DeleteStore)
	}

	// storeReviewRoutes Маршруты для отзывов на магазины
	storeReviewRoutes := r.Group("/store/reviews")
	{
		storeReviewRoutes.GET("/:id", controllers.GetAllStoreReviewsByStoreID)
		storeReviewRoutes.POST("/:id", middlewares.CheckUserAuthentication, controllers.CreateStoreReview)
		storeReviewRoutes.PUT("/:id", middlewares.CheckUserAuthentication, controllers.UpdateStoreReview)
	}

	r.GET("/store/review/:id", controllers.GetStoreReviewByID)
	r.GET("/hash-password", middlewares.CheckSecretKey, controllers.HashPassword)
	r.DELETE("/store/review/:id", middlewares.CheckUserAuthentication, controllers.DeleteStoreReview)

	// categoryRoutes Маршруты для категорий на магазине
	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.GET("/", controllers.GetAllCategories)
		categoryRoutes.GET("/:id", controllers.GetCategoryById)
		categoryRoutes.POST("/", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, controllers.CreateCategory)
		categoryRoutes.PUT("/:id", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, controllers.UpdateCategory)
		categoryRoutes.DELETE("/:id", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, controllers.DeleteCategory)
	}

	// orderStatusGroup Маршруты для статусов заказов
	orderStatusGroup := r.Group("/order-status")
	{
		orderStatusGroup.GET("/", controllers.GetAllOrderStatuses)
		orderStatusGroup.GET("/:id", controllers.GetOrderStatusByID)
		orderStatusGroup.POST("/", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, controllers.CreateOrderStatus)
		orderStatusGroup.PUT("/:id", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, controllers.UpdateOrderStatus)
		orderStatusGroup.DELETE("/:id", middlewares.CheckUserAuthentication, middlewares.CheckAdmin, controllers.DeleteOrderStatus)
	}

	// Обработчик статусов заказов по имени
	r.GET("/order-status/name/:name", controllers.GetOrderStatusByName)

	productGroup := r.Group("/product")
	{
		productGroup.GET("/", controllers.GetAllProducts)
		productGroup.GET("/:id", controllers.GetProductByID)
		productGroup.POST("/:store_id", middlewares.CheckUserAuthentication, controllers.CreateProduct)
		productGroup.PUT("/:id", middlewares.CheckUserAuthentication, controllers.UpdateProduct)
		productGroup.DELETE("/:id", middlewares.CheckUserAuthentication, controllers.DeleteProduct)
	}

	addressGroup := r.Group("/address", middlewares.CheckUserAuthentication)
	{
		addressGroup.GET("/", controllers.GetAddressesByUserID)
		addressGroup.GET("/:id", controllers.GetAddressByID)
		addressGroup.POST("/", controllers.CreateAddress)
		addressGroup.PUT("/:id", controllers.UpdateAddress)
		addressGroup.DELETE("/:id", controllers.DeleteAddress)
	}

	accountGroup := r.Group("/accounts", middlewares.CheckUserAuthentication)
	{
		accountGroup.GET("/", controllers.GetAccountsByUserID)
		accountGroup.GET("/:id", controllers.GetAccountByID)
		accountGroup.POST("/", controllers.CreateAccount)
		accountGroup.PUT("/:id", controllers.UpdateAccount)
		accountGroup.PUT("/fill", controllers.FillAccountBalance)
		accountGroup.DELETE("/:id", controllers.DeleteAccount)
	}

	featuredProductGroup := r.Group("/products/featured", middlewares.CheckUserAuthentication)
	{
		featuredProductGroup.GET("/", controllers.GetFeaturedProducts)
		featuredProductGroup.GET("/:id", controllers.GetFeaturedProductByID)
		featuredProductGroup.POST("/:id", controllers.CreateFeaturedProduct)
		featuredProductGroup.DELETE("/:id", controllers.DeleteFeaturedProduct)
	}

	return r
}
