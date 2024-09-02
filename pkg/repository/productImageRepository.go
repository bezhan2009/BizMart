package repository

import (
	"BizMart/db"
	"BizMart/logger"
	"BizMart/models"
)

func GetProductImageRepository(productID uint) (productImage models.ProductImage, err error) {
	if err = db.GetDBConn().Where("product_id = ?", productID).First(&productImage).Error; err != nil {
		logger.Error.Printf("[repository.GetProductImageRepository] error getting product image repository %v", err)
		return productImage, err
	}

	return productImage, nil
}

func CreateProductImage(productImage models.ProductImage) (err error) {
	if err = db.GetDBConn().Create(&productImage).Error; err != nil {
		logger.Error.Printf("[repository.CreateProductImage] error creating Images for product: %v\n]", err)
		return err
	}

	return nil
}

func UpdateProductImage(productImage models.ProductImage) (err error) {
	if err = db.GetDBConn().Save(&productImage).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProductImage] error updating Images for product: %v\n]", err)
		return err
	}

	return nil
}

func DeleteProductImage(productImageID uint) (err error) {
	if err = db.GetDBConn().Delete(&models.ProductImage{}, productImageID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProductImage] error deleting Images for product: %v\n]", err)
		return err
	}

	return nil
}
