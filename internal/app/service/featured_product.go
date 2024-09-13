package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
)

func AddToFeaturedProducts(featuredProductData models.FeaturedProduct, userID, productID uint) error {
	featuredProductData.UserID = userID
	featuredProductData.ProductID = productID
	if err := repository.CreateFeaturedProduct(featuredProductData); err != nil {
		return err
	}

	return nil
}
