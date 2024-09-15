package repository

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
)

func GetAllProductReviews(productID uint) ([]models.Review, error) {
	var reviews []models.Review
	if err := db.GetDBConn().Model(models.Review{}).Where("product_id = ?", productID).Find(&reviews).Error; err != nil {
		logger.Error.Printf("[repository.GetAllProductReview] Error getting product reviews: %v", err)
		return nil, TranslateGormError(err)
	}

	return reviews, nil
}

func GetProductReviewByID(ProductReviewID uint) (models.Review, error) {
	var review models.Review
	if err := db.GetDBConn().Model(models.Review{}).Where("id = ?", ProductReviewID).First(&review).Error; err != nil {
		logger.Error.Printf("[repository.GetProductReviewByID] Error getting product reviews: %v", err)
		return review, TranslateGormError(err)
	}

	return review, nil
}

func GetProductByUserAndProductID(userID, productID uint) (models.Review, error) {
	var review models.Review
	if err := db.GetDBConn().Model(models.Review{}).Where("user_id = ? AND product_id = ?", userID, productID).First(&review).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByUserAndProductID] Error getting product reviews: %v", err)
		return review, TranslateGormError(err)
	}

	return review, nil
}

func CreateProductReview(review models.Review) error {
	if err := db.GetDBConn().Model(models.Review{}).Create(&review).Error; err != nil {
		logger.Error.Printf("[repository.CreateProductReview] Error creating review: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func UpdateProductReview(review models.Review) error {
	var reviewData models.Review
	if err := db.GetDBConn().Model(models.Review{}).Where("id = ?", review.ID).Find(&reviewData).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProductReview] Error updating review: %v", err)
		return TranslateGormError(err)
	}

	reviewData.Rating = review.Rating
	reviewData.Title = review.Title
	reviewData.Content = review.Content

	if err := db.GetDBConn().Model(models.Review{}).Save(&reviewData).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProductReview] Error updating review: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func DeleteProductReview(productReview models.Review) error {
	if err := db.GetDBConn().Delete(&productReview).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProductReview] Error deleting review: %v", err)
		return TranslateGormError(err)
	}

	return nil
}
