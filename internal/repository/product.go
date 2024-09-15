package repository

import (
	models2 "BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
)

// GetProductByID retrieves a product by its ID
func GetProductByID(productID uint) (models2.Product, error) {
	var product models2.Product
	if err := db.GetDBConn().Preload("Store").Preload("Category").Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error getting product: %v\n", err)
		return product, TranslateGormError(err)
	}

	// Увеличиваем количество просмотров
	product.Views += 1

	// Сохраняем изменения в базе данных
	if err := db.GetDBConn().Save(&product).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error updating product views: %v\n", err)
		return product, TranslateGormError(err)
	}

	return product, nil
}

// DeleteProductByID marks a product as deleted
func DeleteProductByID(productID uint) error {
	// Fetch the existing product
	var product models2.Product
	if err := db.GetDBConn().Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProduct] Error finding product: %v\n", err)
		return TranslateGormError(err)
	}

	if err := db.GetDBConn().Delete(&product).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProduct] Error deleting product: %v\n", err)
		return TranslateGormError(err)
	}

	return nil
}

func DeleteProductImagesByProductID(productID uint) error {
	if err := db.GetDBConn().Where("product_id = ?", productID).Delete(&models2.ProductImage{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProductImagesByProductID] Error deleting product images: %v\n", err)
		return TranslateGormError(err)
	}
	return nil
}

// GetAllProducts retrieves all products filtered by category, price range, etc., and includes associated ProductImage data.
func GetAllProducts(minPrice, maxPrice float64, categoryID uint, productName string, storeID uint) ([]models2.Product, error) {
	var products []models2.Product

	// Start building the query
	query := db.GetDBConn().
		Table("productapp_product").
		Select("productapp_product.*, STRING_AGG(productapp_productimage.image, ',') AS product_images").
		Where("productapp_product.amount > 0").
		Joins("LEFT JOIN productapp_productimage ON productapp_productimage.product_id = productapp_product.id").
		Group("productapp_product.id")

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
		query = query.Where(
			db.GetDBConn().Where("productapp_product.title LIKE ?", "%"+productName+"%").
				Or("productapp_product.description LIKE ?", "%"+productName+"%"),
		)
	}

	// Execute the query
	if err := query.Find(&products).Error; err != nil {
		logger.Error.Printf("[repository.GetAllProducts] Error retrieving products: %v\n", err)
		return nil, TranslateGormError(err)
	}

	return products, nil
}

func CreateProductWithImages(product *models2.Product, images []models2.ProductImage) error {
	// Создаем продукт в базе данных
	if err := db.GetDBConn().Create(product).Error; err != nil {
		return TranslateGormError(err)
	}

	// Присваиваем ID продукта для всех изображений
	for i := range images {
		images[i].ProductID = product.ID
	}

	// Сохраняем все изображения в базе данных
	if err := db.GetDBConn().Create(&images).Error; err != nil {
		return TranslateGormError(err)
	}

	return nil
}

func UpdateProductWithImages(product *models2.Product, images []models2.ProductImage) error {
	// Обновляем продукт в базе данных
	if err := db.GetDBConn().Save(product).Error; err != nil {
		return TranslateGormError(err)
	}

	// Удаляем старые изображения
	if err := db.GetDBConn().Where("product_id = ?", product.ID).Delete(&models2.ProductImage{}).Error; err != nil {
		return TranslateGormError(err)
	}

	// Добавляем новые изображения
	for i := range images {
		images[i].ProductID = product.ID
	}

	// Сохраняем новые изображения
	if err := db.GetDBConn().Create(&images).Error; err != nil {
		return TranslateGormError(err)
	}

	return nil
}
