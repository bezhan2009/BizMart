package Product

import (
	"BizMart/errs"
	"BizMart/models"
	"BizMart/pkg/controllers/handlers"
	"BizMart/pkg/repository/product"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllProducts(c *gin.Context) {
	minPriceStr := c.Query("minPrice")
	maxPriceStr := c.Query("maxPrice")
	category := c.Query("category")
	productName := c.Query("productName")
	store := c.Query("store")

	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		handlers.HandleError(c, errs.ErrInvalidMinPrice)
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		handlers.HandleError(c, errs.ErrInvalidMaxPrice)
		return
	}

	categoryId, err := strconv.Atoi(category)
	if err != nil {
		handlers.HandleError(c, errs.ErrInvalidCategory)
		return
	}

	storeId, err := strconv.Atoi(store)
	if err != nil {
		handlers.HandleError(c, errs.ErrInvalidCategory)
		return
	}

	var products []models.Product

	products, err = product.GetAllProducts(minPrice, maxPrice, uint(categoryId), productName, uint(storeId))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching products"})
		return
	}

	c.JSON(200, gin.H{"products": products})
}
