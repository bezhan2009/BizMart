package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/controllers/middlewares"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
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

	if minPriceStr == "" && maxPriceStr == "" && category == "" && store == "" {
		var products []models.Product

		products, err := repository.GetAllProducts(0, 0, uint(0), productName, uint(0))
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching products"})
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
		c.JSON(500, gin.H{"error": "Error fetching products"})
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

	c.JSON(http.StatusOK, gin.H{"getProductByID": getProductByID})
}

func CreateProduct(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductID)
		return
	}

	var productData models.Product
	if err := c.ShouldBind(&productData); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Получаем данные о магазине
	productData.Store, err = repository.GetStoreByID(uint(storeID))
	if err != nil {
		HandleError(c, errs.ErrProductNotFound)
		return
	}

	// Получаем ID пользователя
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Проверка прав на создание продукта
	if userID != productData.Store.OwnerID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	// Извлекаем поле product_image, ожидая массив строк (ссылки на изображения)
	var productImages []string
	if err := c.BindJSON(&productImages); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Создаем массив структур ProductImage на основе пришедших данных
	var images []models.ProductImage
	for _, image := range productImages {
		images = append(images, models.ProductImage{
			ProductID: productData.ID, // ID продукта будет присвоен после создания продукта
			Image:     image,
		})
	}

	// Сохраняем продукт и изображения
	if err := repository.CreateProductWithImages(&productData, images, userID); err != nil {
		HandleError(c, err)
		return
	}

	// Ответ клиенту
	c.JSON(http.StatusOK, gin.H{
		"message": "Product and images successfully created",
		"product": productData,
	})
}
