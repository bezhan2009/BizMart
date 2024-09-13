package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/controllers/middlewares"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetFeaturedProducts godoc
// @Summary      Get all featured products
// @Description  Retrieves a list of all featured products for the authenticated user
// @Tags         featured_products
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.FeaturedProduct
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       //products/featured [get]
func GetFeaturedProducts(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	featuredProducts, err := repository.GetAllFeaturedProducts(userID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrFeaturedProductNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"featured_products": featuredProducts})
}

// GetFeaturedProductByID godoc
// @Summary      Get featured product by ID
// @Description  Retrieves a specific featured product by its ID for the authenticated user
// @Tags         featured_products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Featured Product ID"
// @Success      200  {object}  models.FeaturedProduct
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      403  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       //products/featured/{id} [get]
func GetFeaturedProductByID(c *gin.Context) {
	featuredProductIdStr := c.Param("id")
	featuredProductId, err := strconv.Atoi(featuredProductIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidFeaturedProductID)
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	featuredProduct, err := repository.GetFeaturedProductByID(uint(featuredProductId))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrFeaturedProductNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if userID != featuredProduct.UserID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	c.JSON(http.StatusOK, gin.H{"featured_product": featuredProduct})
}

// CreateFeaturedProduct godoc
// @Summary      Create featured product
// @Description  Adds a new product to the user's featured products list
// @Tags         featured_products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Param        product  body  models.FeaturedProductsRequest  true  "Featured Product Data"
// @Success      200  {object}  models.DefaultResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       //products/featured/{id} [post]
func CreateFeaturedProduct(c *gin.Context) {
	ProductIdStr := c.Param("id")
	ProductId, err := strconv.Atoi(ProductIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidFeaturedProductID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var featuredProduct models.FeaturedProduct
	if err := c.Bind(&featuredProduct); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.AddToFeaturedProducts(featuredProduct, userID, uint(ProductId)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "featured product added successfully"})
}

// DeleteFeaturedProduct godoc
// @Summary      Delete featured product
// @Description  Deletes a specific featured product by its ID
// @Tags         featured_products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Featured Product ID"
// @Success      200  {object}  models.DefaultResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      403  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /products/featured/{id} [delete]
func DeleteFeaturedProduct(c *gin.Context) {
	featuredProductIdStr := c.Param("id")
	featuredProductId, err := strconv.Atoi(featuredProductIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidFeaturedProductID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	featuredProduct, err := repository.GetFeaturedProductByID(uint(featuredProductId))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrFeaturedProductNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if userID != featuredProduct.UserID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	if err := repository.DeleteFeaturedProduct(featuredProduct); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrFeaturedProductNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "featured product deleted successfully"})
}
