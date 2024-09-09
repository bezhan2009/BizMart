package repository

import (
	models2 "BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// CreateProduct creates a new product in the store
func CreateProduct(product *models2.Product, userID uint, productImage models2.ProductImage) error {
	// Validate that the user owns the store
	var store models2.Store
	if err := db.GetDBConn().Where("id = ? AND owner_id = ?", product.StoreID, userID).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("unauthorized action: you do not own this store")
		}
		return err
	}

	// Create the product
	if err := db.GetDBConn().Create(&product).Error; err != nil {
		logger.Error.Printf("[repository.CreateProduct] Error creating product: %v\n", err)
		return err
	}

	err := CreateProductImage(productImage)
	if err != nil {
		return err
	}

	return nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(productID uint) (models2.Product, error) {
	var product models2.Product
	if err := db.GetDBConn().Preload("Store").Preload("Category").Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error getting product: %v\n", err)
		return product, err
	}

	// Увеличиваем количество просмотров
	product.Views += 1

	// Сохраняем изменения в базе данных
	if err := db.GetDBConn().Save(&product).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error updating product views: %v\n", err)
		return product, err
	}

	return product, nil
}

func UpdateProduct(productID uint, updatedProduct *models2.Product, updatedImages []models2.ProductImage) error {
	// Fetch the existing product
	var product models2.Product
	if err := db.GetDBConn().Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProduct] Error finding product: %v\n", err)
		return err
	}

	// Check if the category exists
	var category models2.Category
	if err := db.GetDBConn().Where("id = ?", updatedProduct.CategoryID).First(&category).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProduct] Category not found for ID: %v, Error: %v\n", updatedProduct.CategoryID, err)
		return fmt.Errorf("category with ID %d not found", updatedProduct.CategoryID)
	}

	// Update the product details
	product.Title = updatedProduct.Title
	product.Description = updatedProduct.Description
	product.Price = updatedProduct.Price
	product.Amount = updatedProduct.Amount
	product.CategoryID = updatedProduct.CategoryID

	//if updatedProduct.DefaultAccountID != 0 {
	//	product.DefaultAccountID = updatedProduct.DefaultAccountID
	//}

	// Start transaction to ensure atomicity
	tx := db.GetDBConn().Begin()

	// Save updated product details
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		logger.Error.Printf("[repository.UpdateProduct] Error updating product: %v\n", err)
		return err
	}

	// Remove old product images
	if err := tx.Where("product_id = ?", productID).Delete(&models2.ProductImage{}).Error; err != nil {
		tx.Rollback()
		logger.Error.Printf("[repository.UpdateProduct] Error deleting old product images: %v\n", err)
		return err
	}

	// Assign the updated product ID to the new images and save them
	for i := range updatedImages {
		updatedImages[i].ProductID = productID
	}

	// Save the updated product images
	if err := tx.Create(&updatedImages).Error; err != nil {
		tx.Rollback()
		logger.Error.Printf("[repository.UpdateProduct] Error updating product images: %v\n", err)
		return err
	}

	// Commit the transaction
	tx.Commit()

	return nil
}

// DeleteProductByID marks a product as deleted
func DeleteProductByID(productID uint) error {
	// Fetch the existing product
	var product models2.Product
	if err := db.GetDBConn().Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProduct] Error finding product: %v\n", err)
		return err
	}

	if err := db.GetDBConn().Delete(&product).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProduct] Error deleting product: %v\n", err)
		return err
	}

	return nil
}

func DeleteProductImagesByProductID(productID uint) error {
	if err := db.GetDBConn().Where("product_id = ?", productID).Delete(&models2.ProductImage{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProductImagesByProductID] Error deleting product images: %v\n", err)
		return err
	}
	return nil
}

// GetAllProducts retrieves all products filtered by category, price range, etc., and includes associated ProductImage data.
func GetAllProducts(minPrice, maxPrice float64, categoryID uint, productName string, storeID uint) ([]models2.Product, error) {
	var products []models2.Product

	// Start building the query
	query := db.GetDBConn().
		Table("productapp_product").
		Select("productapp_product.*, productapp_productimage.image AS product_image").
		Joins("LEFT JOIN productapp_productimage ON productapp_productimage.product_id = productapp_product.id").
		Where("productapp_product.is_deleted = ?", false)

	// Apply filters
	if minPrice > 0 {
		query = query.Where("productapp_product.price >= ?", minPrice)
	}
	if maxPrice > 0 {
		query = query.Where("productapp_product.price <= ?", maxPrice)
	}
	if categoryID > 0 {
		query = query.Where("productapp_product.category_id = ?", categoryID)
	}
	if storeID > 0 {
		query = query.Where("productapp_product.store_id = ?", storeID)
	}
	if productName != "" {
		query = query.Where("productapp_product.title LIKE ?", "%"+productName+"%")
	}

	// Execute the query
	if err := query.Find(&products).Error; err != nil {
		logger.Error.Printf("[repository.GetAllProducts] Error retrieving products: %v\n", err)
		return nil, err
	}

	return products, nil
}

func CreateProductWithImages(product *models2.Product, images []models2.ProductImage) error {
	// Создаем продукт в базе данных
	if err := db.GetDBConn().Create(product).Error; err != nil {
		return err
	}

	// Присваиваем ID продукта для всех изображений
	for i := range images {
		images[i].ProductID = product.ID
	}

	// Сохраняем все изображения в базе данных
	if err := db.GetDBConn().Create(&images).Error; err != nil {
		return err
	}

	return nil
}

func UpdateProductWithImages(product *models2.Product, images []models2.ProductImage) error {
	// Обновляем продукт в базе данных
	if err := db.GetDBConn().Save(product).Error; err != nil {
		return err
	}

	// Удаляем старые изображения
	if err := db.GetDBConn().Where("product_id = ?", product.ID).Delete(&models2.ProductImage{}).Error; err != nil {
		return err
	}

	// Добавляем новые изображения
	for i := range images {
		images[i].ProductID = product.ID
	}

	// Сохраняем новые изображения
	if err := db.GetDBConn().Create(&images).Error; err != nil {
		return err
	}

	return nil
}
