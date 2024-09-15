package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
)

func AddToFeaturedProducts(featuredProductData models.FeaturedProduct, userID, productID uint) error {
	featuredProductData.UserID = userID
	featuredProductData.ProductID = productID
	if _, err := repository.GetProductByID(productID); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrProductNotFound
		}

		return err
	}

	var err error

	if _, err = repository.GetFeaturedProductByUserAndProductId(productID, userID); err != nil {
		if errors.Is(err, errs.ErrProductNotFound) {
			return errs.ErrProductNotFound
		}
	}

	if err == nil {
		return errs.ErrFeaturedProductUniquenessFailed
	}

	if err = repository.CreateFeaturedProduct(featuredProductData); err != nil {
		return err
	}

	return nil
}
