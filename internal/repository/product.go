package repository

import (
	models2 "BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
	"sort"
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
		logger.Error.Printf("[repository.DeleteProductByID] Error finding product: %v\n", err)
		return TranslateGormError(err)
	}

	if err := db.GetDBConn().Delete(&product).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProductByID] Error deleting product: %v\n", err)
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

// GetAllProducts retrieves all products filtered by category, price range, etc.,
// and sorts them by the number of orders and views.
func GetAllProducts(minPrice, maxPrice float64, categoryID uint, productName string, storeID uint) ([]models2.Product, error) {
	var products []models2.Product
	isQuery := false
	query := db.GetDBConn().
		Table("productapp_product").
		Select("productapp_product.*").
		Where("productapp_product.amount > 0")

	if minPrice > 0 {
		isQuery = true
		query = query.Where("productapp_product.price >= ?", minPrice)
	}
	if maxPrice > 0 {
		isQuery = true
		query = query.Where("productapp_product.price <= ?", maxPrice)
	}
	if categoryID > 0 {
		isQuery = true
		query = query.Where("productapp_product.category_id = ?", categoryID)
	}
	if storeID > 0 {
		isQuery = true
		query = query.Where("productapp_product.store_id = ?", storeID)
	}
	if productName != "" {
		isQuery = true
		query = query.Where(
			db.GetDBConn().Where("productapp_product.title LIKE ?", "%"+productName+"%").
				Or("productapp_product.description LIKE ?", "%"+productName+"%"),
		)
	}

	if err := query.Find(&products).Error; err != nil {
		logger.Error.Printf("[repository.GetAllProducts] Error retrieving products: %v\n", err)
		return nil, TranslateGormError(err)
	}
	if !isQuery {
		for i, product := range products {
			numOfOrders, err := GetNumberOfProductOrders(product.ID)
			if err != nil {
				logger.Error.Printf("[repository.GetAllProducts] Error retrieving number of orders: %v\n", err)
				return nil, TranslateGormError(err)
			}
			products[i].Amount = uint(numOfOrders)
		}

		sort.Slice(products, func(i, j int) bool {
			if products[i].Amount == products[j].Amount {
				return products[i].Views > products[j].Views
			}
			return products[i].Amount > products[j].Amount
		})
	}
	return products, nil
}

func CreateProductWithImages(product *models2.Product, images []models2.ProductImage) error {
	// Создаем продукт в базе данных
	if err := db.GetDBConn().Create(product).Error; err != nil {
		logger.Error.Printf("[repository.CreateProductWithImages] error creating product: %v\n", err)
		return TranslateGormError(err)
	}

	// Присваиваем ID продукта для всех изображений
	for i := range images {
		images[i].ProductID = product.ID
	}

	// Сохраняем все изображения в базе данных
	if err := db.GetDBConn().Create(&images).Error; err != nil {
		logger.Error.Printf("[repository.CreateProductWithImages] error creating product images: %v\n", err)
		return TranslateGormError(err)
	}

	return nil
}

func UpdateProductWithImages(product *models2.Product, images []models2.ProductImage) error {
	// Обновляем продукт в базе данных
	if err := db.GetDBConn().Save(product).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProductWithImages] error updating product: %v\n", err)
		return TranslateGormError(err)
	}

	// Удаляем старые изображения
	if err := db.GetDBConn().Where("product_id = ?", product.ID).Delete(&models2.ProductImage{}).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProductWithImages] error deleting product image: %v\n", err)
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

func UpdateProduct(product *models2.Product) error {
	if err := db.GetDBConn().Model(&models2.Product{}).Where("id = ?", product.ID).Updates(map[string]interface{}{
		"amount": product.Amount,
	}).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProduct] Error updating product: %v\n", err)
		return TranslateGormError(err)
	}

	return nil
}

func GetProductByStoreID(storeID uint) ([]models2.Product, error) {
	var products []models2.Product
	if err := db.GetDBConn().Model(&models2.Product{}).Where("store_id = ?", storeID).Find(&products).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByStoreID] Error getting product: %v\n", err)
		return nil, TranslateGormError(err)
	}

	return products, nil
}

func GetProductByStoreIDWithoutFilters(storeID uint) ([]models2.Product, error) {
	var products []models2.Product
	if err := db.GetDBConn().
		Model(&models2.Product{}).
		Preload("Store").
		Where("store_id = ?", storeID).
		Find(&products).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByStoreID] Error getting product: %v\n", err)
		return nil, TranslateGormError(err)
	}

	return products, nil
}
