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
