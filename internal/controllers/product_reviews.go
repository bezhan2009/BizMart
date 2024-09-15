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

// GetAllProductReviews godoc
// @Summary      Get all products reviews
// @Description  Retrieves a list of all products reviews
// @Tags         product_reviews
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  models.Review
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products/reviews/{id} [get]
func GetAllProductReviews(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductID)
		return
	}

	reviews, err := repository.GetAllProductReviews(uint(productId))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_reviews": reviews})
}

// GetProductReviewByID godoc
// @Summary      Get product review By ID
// @Description  Retrieves a review
// @Tags         product_reviews
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product Review ID"
// @Success      200  {object}  models.Review
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products/review/{id} [get]
func GetProductReviewByID(c *gin.Context) {
	productReviewIdStr := c.Param("id")
	productReviewId, err := strconv.Atoi(productReviewIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductReviewID)
		return
	}

	review, err := repository.GetProductReviewByID(uint(productReviewId))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrProductReviewNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_review": review})
}

// CreateProductReview godoc
// @Summary      Create product review
// @Description  Creates a new product review
// @Tags         product_reviews
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Param        product  body  models.ReviewRequest  true  "Product Review Data"
// @Success      200  {object}  models.DefaultResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      403  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products/reviews/{id} [post]
func CreateProductReview(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductID)
		return
	}

	_, err = repository.GetProductByID(uint(productId))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrProductNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	review.UserID = userID
	review.ProductID = uint(productId)

	if err := service.ValidateProductReview(HandleError, review, c); err != nil {
		return
	}

	if err := service.CreateProductReview(review); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product review created successfully"})
}

// UpdateProductReview godoc
// @Summary      Update product review
// @Description  Updates a new product review
// @Tags         product_reviews
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product Review ID"
// @Param        product  body  models.ReviewRequest  true  "Product Review Data"
// @Success      200  {object}  models.DefaultResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      403  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products/reviews/{id} [put]
func UpdateProductReview(c *gin.Context) {
	productReviewIdStr := c.Param("id")
	productReviewId, err := strconv.Atoi(productReviewIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductReviewID)
		return
	}

	reviewData, err := repository.GetProductReviewByID(uint(productReviewId))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrProductReviewNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	if reviewData.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	review.ProductID = reviewData.ProductID
	review.ID = reviewData.ID

	if err := service.ValidateProductReview(HandleError, review, c); err != nil {
		return
	}

	if err = repository.UpdateProductReview(review); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrProductReviewNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product review updated successfully"})
}

// DeleteProductReview godoc
// @Summary      Delete product review
// @Description  Deletes a new product review
// @Tags         product_reviews
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product Review ID"
// @Success      200  {object}  models.DefaultResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      403  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /products/reviews/{id} [delete]
func DeleteProductReview(c *gin.Context) {
	productReviewIdStr := c.Param("id")
	productReviewId, err := strconv.Atoi(productReviewIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidProductReviewID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	review, err := repository.GetProductReviewByID(uint(productReviewId))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrProductReviewNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if review.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	if err := repository.DeleteProductReview(review); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrProductReviewNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product review deleted successfully"})
}
