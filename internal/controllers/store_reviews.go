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

// GetAllStoreReviewsByStoreID godoc
// @Summary Get all reviews for a store
// @Description Fetches all reviews for a specific store by its ID.
// @Tags store reviews
// @Accept  json
// @Produce  json
// @Param id path int true "Store ID"
// @Success 200 {object} models.Store "Returns a list of reviews for the store"
// @Failure 400 {object} models.ErrorResponse
// @Router /store/reviews/{id} [get]
func GetAllStoreReviewsByStoreID(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidStoreID)
		return
	}

	storeReviews, err := repository.GetAllStoreReviews(uint(storeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"store_reviews": storeReviews})
}

// GetStoreReviewByID godoc
// @Summary Get a specific store review by its ID
// @Description Retrieves a specific review by its ID.
// @Tags store reviews
// @Accept  json
// @Produce  json
// @Param id path int true "Store Review ID"
// @Success 200 {object} models.Store "Returns the store review"
// @Failure 404 {object} models.ErrorResponse
// @Router /store/reviews/{id} [get]
func GetStoreReviewByID(c *gin.Context) {
	storeReviewIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeReviewIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidStoreReviewID)
		return
	}

	storeReview, err := repository.GetStoreReviewByID(uint(storeID))
	if err != nil {
		HandleError(c, errs.ErrStoreReviewNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"store_review": storeReview})
}

// CreateStoreReview godoc
// @Summary Create a new store review
// @Description Creates a new review for a store.
// @Tags store reviews
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Store ID"
// @Param review body models.StoreReviewRequest true "Store Review data"
// @Success 200 {object} models.DefaultResponse "Returns success message"
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /store/reviews/{id} [post]
func CreateStoreReview(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidStoreID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	var storeReview models.StoreReview
	if err := c.ShouldBindJSON(&storeReview); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	storeReviewAndStoreIdData, err := repository.GetStoreReviewByStoreIdAndUserId(userID, uint(storeID))
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		HandleError(c, err)
		return
	}

	if storeReviewAndStoreIdData != nil {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	storeReview.UserID = userID
	storeReview.StoreID = uint(storeID)

	if err = service.ValidateStoreReview(HandleError, storeReview, c); err != nil {
		return
	}

	if err := repository.CreateStoreReview(&storeReview); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store review created successfully"})
}

// UpdateStoreReview godoc
// @Summary Update an existing store review
// @Description Updates the details of an existing store review.
// @Tags store reviews
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Store Review ID"
// @Param review body models.StoreReviewRequest true "Updated store review data"
// @Success 200 {object} models.DefaultResponse "Returns success message"
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /store/reviews/{id} [put]
func UpdateStoreReview(c *gin.Context) {
	storeReviewIDStr := c.Param("id")
	storeReviewID, err := strconv.Atoi(storeReviewIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidStoreID)
		return
	}

	storeReviewIdData, err := repository.GetStoreReviewByID(uint(storeReviewID))
	if err != nil {
		HandleError(c, err)
		return
	}

	var storeReview models.StoreReview
	if err = c.ShouldBindJSON(&storeReview); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	if storeReviewIdData.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	storeReview.UserID = userID
	storeReview.StoreID = storeReviewIdData.StoreID

	if err = service.ValidateStoreReview(HandleError, storeReview, c); err != nil {
		return
	}

	err = repository.UpdateStoreReview(uint(storeReviewID), &storeReview)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrStoreReviewNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store review updated successfully"})
}

// DeleteStoreReview godoc
// @Summary Delete a store review
// @Description Deletes a specific review by its ID.
// @Tags store reviews
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Store Review ID"
// @Success 200 {object} models.DefaultResponse "Returns success message"
// @Failure 404 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /store/reviews/{id} [delete]
func DeleteStoreReview(c *gin.Context) {
	storeReviewIDStr := c.Param("id")
	storeReviewID, err := strconv.Atoi(storeReviewIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidStoreID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	storeReview, err := repository.GetStoreReviewByID(uint(storeReviewID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrStoreReviewNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if storeReview.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	err = repository.DeleteStoreReview(uint(storeReviewID))
	if err != nil {
		HandleError(c, errs.ErrDeleteFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store review deleted successfully"})
}
