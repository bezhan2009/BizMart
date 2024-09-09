package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/controllers/middlewares"
	"BizMart/internal/repository"
	"BizMart/pkg/db"
	"BizMart/pkg/errs"
	"BizMart/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllProducts(c *gin.Context) {
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")
	category := c.Query("category")
	productName := c.Query("product_name")
	store := c.Query("store")

	if minPriceStr == "" && maxPriceStr == "" && category == "" && store == "" && productName == "" {
		var products []models.Product

		products, err := repository.GetAllProducts(0, 0, uint(0), productName, uint(0))
		if err != nil {
			HandleError(c, errs.ErrFetchingProducts)
			return
		}

		if len(products) == 0 {
			HandleError(c, errs.ErrNoProductsFound)
			return
		}

		c.JSON(200, gin.H{"products": products})
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

	if minPrice > maxPrice {
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
		HandleError(c, errs.ErrFetchingProducts)
		return
	}

	if len(products) == 0 {
		HandleError(c, errs.ErrNoProductsFound)
		return
	}

	c.JSON(200, gin.H{"products": products})
}

func GetProductByID(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductID)
		return
	}

	getProductByID, err := repository.GetProductByID(uint(productId))
	if err != nil {
		HandleError(c, errs.ErrProductNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": getProductByID})
}

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

	err = service.ValidateProduct(HandleError, productData, c)
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
		HandleError(c, errs.ErrPermissionDenied)
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
		"product": productData,
	})
}

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
		HandleError(c, errs.ErrProductNotFound)
		return
	}

	// Получаем данные из запроса
	var updatedProductData models.Product
	if err := c.ShouldBindJSON(&updatedProductData); err != nil {
		logger.Error.Printf("[controllers.UpdateProduct] Error updating product: %s", err.Error())
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.ValidateProduct(HandleError, productData, c)
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
		"product": productData,
	})
}

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
		HandleError(c, errs.ErrProductNotFound)
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
