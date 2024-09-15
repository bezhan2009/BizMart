package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
)

func ValidateProductReview(HandleError func(ctx *gin.Context, err error), productReviewData models.Review, c *gin.Context) error {
	if productReviewData.ProductID == 0 {
		HandleError(c, errs.ErrInvalidProductID)
		return errs.ErrInvalidProductID
	}

	if productReviewData.Rating < 0 || productReviewData.Rating > 5 {
		HandleError(c, errs.ErrInvalidRating)
		return errs.ErrInvalidRating
	}

	if len(productReviewData.Title) < 5 {
		HandleError(c, errs.ErrInvalidTitle)
		return errs.ErrInvalidTitle
	}

	if len(productReviewData.Content) < 5 {
		HandleError(c, errs.ErrInvalidContent)
		return errs.ErrInvalidContent
	}

	return nil
}

func CreateProductReview(review models.Review) error {
	_, err := repository.GetProductByUserAndProductID(review.UserID, review.ProductID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return repository.CreateProductReview(review)
		}
		return err
	}

	return errs.ErrPermissionDenied
}
