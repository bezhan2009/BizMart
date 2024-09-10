package service

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/errs"
	"github.com/gin-gonic/gin"
)

func ValidateStoreReview(HandleError func(ctx *gin.Context, err error), storeReviewData models.StoreReview, c *gin.Context) error {
	if len(storeReviewData.Comment) <= 10 {
		HandleError(c, errs.ErrInvalidComment)
		return errs.ErrInvalidComment
	}

	if storeReviewData.Rating <= 0 || storeReviewData.Rating > 5 {
		HandleError(c, errs.ErrInvalidRating)
		return errs.ErrInvalidRating
	}

	return nil
}
