package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/controllers/middlewares"
	"BizMart/internal/jobs"
	"BizMart/internal/repository"
	"BizMart/pkg/db"
	"BizMart/pkg/errs"
	"BizMart/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllProducts godoc
// @Summary Get all products
// @Description Fetches all products with optional filtering by price, category, product name, and store.
// @Tags products
// @Accept  json
// @Produce  json
// @Param min_price query number false "Minimum price filter"
// @Param max_price query number false "Maximum price filter"
// @Param category query int false "Category ID"
// @Param product_name query string false "Product name"
// @Param store query int false "Store ID"
// @Success 200 {object} models.ProductResponse "Returns a list of products"
// @Failure 400 {object} models.ErrorResponse
// @Router /product [get]
func GetAllProducts(c *gin.Context) {
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")
	category := c.Query("category")
	productName := c.Query("product_name")
	store := c.Query("store")

	if minPriceStr == "" && maxPriceStr == "" && category == "" && store == "" && productName == "" {
		products, err := jobs.GetCachedProducts()
		if err != nil {
			HandleError(c, err)
		}

		c.JSON(http.StatusOK, gin.H{"products": products})
		return
	}

	var minPrice, maxPrice float64
	var err error

	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			HandleError(c, errs.ErrInvalidMinPrice)
			return
		}
	}

	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			HandleError(c, errs.ErrInvalidMaxPrice)
			return
		}
	}

	if minPrice > maxPrice && maxPriceStr != "" {
		HandleError(c, errs.ErrInvalidMinPrice)
		return
	}

	var categoryId int

	if category != "" {
		categoryId, err = strconv.Atoi(category)
		if err != nil {
			HandleError(c, errs.ErrInvalidCategory)
			return
		}
	}

	var storeId int

	if store != "" {
		storeId, err = strconv.Atoi(store)
		if err != nil {
			HandleError(c, errs.ErrInvalidCategory)
			return
		}
	}

	var products []models.Product

	products, err = repository.GetAllProducts(minPrice, maxPrice, uint(categoryId), productName, uint(storeId))
	if err != nil {
		HandleError(c, err)
		return
	}

	if len(products) == 0 {
		HandleError(c, errs.WarningNoProductsFound)
		return
	}

	c.JSON(200, gin.H{"products": products})
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Retrieves a product by its ID along with the number of orders.
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} models.ProductResponse "Returns the product and order count"
// @Failure 404 {object} models.ErrorResponse
// @Router /product/{id} [get]
func GetProductByID(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductID)
		return
	}

	getProductByID, err := repository.GetProductByID(uint(productId))
	if err != nil {
		c.JSON(404, gin.H{"message": errs.ErrNoProductFound.Error()})
		return
	}

	ordersNum, err := repository.GetNumberOfProductOrders(uint(productId))
	if err != nil {
		HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"product": getProductByID,
		"orders":  ordersNum,
	})
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Adds a new product to a specific store, with validation and ownership checks.
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param store_id path int true "Store ID"
// @Param product body models.ProductRequest true "Product data"
// @Success 200 {object} models.DefaultResponse "Returns success message and created product"
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /stores/{store_id}/products [post]
func CreateProduct(c *gin.Context) {
	storeIDStr := c.Param("store_id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidStoreID)
		return
	}

	var productData models.Product
	if err := c.ShouldBindJSON(&productData); err != nil {
		logger.Error.Printf("[controllers.CreateProduct] Error creating new product: %s", err.Error())
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.ValidateProduct(HandleError, productData, c, false)
	if err != nil {
		return
	}

	// Получаем данные о магазине
	productData.Store, err = repository.GetStoreByID(uint(storeID))
	if err != nil {
		HandleError(c, errs.ErrStoreNotFound)
		return
	}

	// Получаем ID пользователя
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	// Проверка прав на создание продукта
	if userID != productData.Store.OwnerID {
		HandleError(c, errs.ErrStoreNotFound)
		return
	}

	// Создаем массив структур ProductImage на основе данных из productData
	var images []models.ProductImage
	for _, image := range productData.ProductImageList {
		images = append(images, models.ProductImage{
			ProductID: productData.ID,
			Image:     image,
		})
	}

	// Сохраняем продукт и изображения
	if err := repository.CreateProductWithImages(&productData, images); err != nil {
		HandleError(c, err)
		return
	}

	// Ответ клиенту
	c.JSON(http.StatusOK, gin.H{
		"message": "Product and images successfully created",
	})
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Updates the details of a product including title, description, price, and images.
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body models.ProductRequest true "Updated product data"
// @Success 200 {object} models.DefaultResponse "Returns success message and updated product"
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductID)
		return
	}

	// Получаем текущие данные продукта из базы данных
	productData, err := repository.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(404, gin.H{"message": errs.ErrNoProductFound.Error()})
		return
	}

	// Получаем данные из запроса
	var updatedProductData models.Product
	if err := c.ShouldBindJSON(&updatedProductData); err != nil {
		logger.Error.Printf("[controllers.UpdateProduct] Error updating product: %s", err.Error())
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.ValidateProduct(HandleError, productData, c, true)
	if err != nil {
		return
	}

	// Получаем ID пользователя
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Проверяем, является ли пользователь владельцем продукта
	if userID != productData.Store.OwnerID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	// Обновляем данные продукта
	productData.Title = updatedProductData.Title
	productData.Description = updatedProductData.Description
	productData.Price = updatedProductData.Price
	productData.Amount = updatedProductData.Amount
	productData.CategoryID = updatedProductData.CategoryID

	// Обновляем Store только в случае необходимости, если это допускается

	// Обновляем изображения
	var updatedImages []models.ProductImage
	for _, image := range updatedProductData.ProductImageList {
		updatedImages = append(updatedImages, models.ProductImage{
			ProductID: productData.ID,
			Image:     image,
		})
	}

	// Сохраняем изменения в базе данных
	if err := repository.UpdateProductWithImages(&productData, updatedImages); err != nil {
		HandleError(c, err)
		return
	}

	// Ответ клиенту
	c.JSON(http.StatusOK, gin.H{
		"message": "Product and images successfully updated",
	})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Deletes a product and its associated images from the database.
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} models.DefaultResponse "Returns success message"
// @Failure 404 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductID)
		return
	}

	// Получаем продукт по ID
	product, err := repository.GetProductByID(uint(productId))
	if err != nil {
		c.JSON(404, gin.H{"message": errs.ErrNoProductFound.Error()})
		return
	}

	// Получаем информацию о магазине
	store, err := repository.GetStoreByID(uint(product.StoreID))
	if err != nil {
		HandleError(c, errs.ErrStoreNotFound)
		return
	}

	// Получаем ID пользователя
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Проверяем права пользователя на удаление продукта
	if store.OwnerID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	// Начинаем транзакцию для удаления продукта и связанных изображений
	tx := db.GetDBConn().Begin()

	// Удаляем изображения продукта
	if err := repository.DeleteProductImagesByProductID(uint(productId)); err != nil {
		tx.Rollback() // Откат транзакции в случае ошибки
		HandleError(c, errs.ErrDeleteFailed)
		return
	}

	// Удаляем сам продукт
	if err := repository.DeleteProductByID(uint(productId)); err != nil {
		tx.Rollback() // Откат транзакции в случае ошибки
		HandleError(c, errs.ErrDeleteFailed)
		return
	}

	// Фиксируем транзакцию после успешного удаления
	tx.Commit()

	// Ответ клиенту об успешном удалении
	c.JSON(http.StatusOK, gin.H{
		"message": "Product and images successfully deleted",
	})
}
