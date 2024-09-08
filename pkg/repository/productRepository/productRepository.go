package productRepository

import (
	"BizMart/db"
	"BizMart/logger"
	"BizMart/models"
	"errors"
	"gorm.io/gorm"
)

// CreateProduct creates a new product in the store
func CreateProduct(product *models.Product, userID uint, productImage models.ProductImage) error {
	// Validate that the user owns the store
	var store models.Store
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
func GetProductByID(productID uint) (models.Product, error) {
	var product models.Product
	if err := db.GetDBConn().Preload("Store").Preload("Category").Preload("DefaultAccount").Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.GetProductByID] Error getting product: %v\n", err)
		return product, err
	}
	return product, nil
}

// UpdateProduct updates an existing product in the store
func UpdateProduct(productID uint, updatedProduct *models.Product, userID uint) error {
	// Fetch the existing product
	var product models.Product
	if err := db.GetDBConn().Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProduct] Error finding product: %v\n", err)
		return err
	}

	// Validate that the user owns the store associated with the product
	if product.Store.OwnerID != userID {
		return errors.New("unauthorized action: you do not own this store")
	}

	// Update the product details
	product.Title = updatedProduct.Title
	product.Description = updatedProduct.Description
	product.Price = updatedProduct.Price
	product.Amount = updatedProduct.Amount
	product.CategoryID = updatedProduct.CategoryID
	product.DefaultAccountID = updatedProduct.DefaultAccountID

	if err := db.GetDBConn().Save(&product).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProduct] Error updating product: %v\n", err)
		return err
	}

	return nil
}

// DeleteProduct marks a product as deleted
func DeleteProduct(productID uint, userID uint) error {
	// Fetch the existing product
	var product models.Product
	if err := db.GetDBConn().Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProduct] Error finding product: %v\n", err)
		return err
	}

	// Validate that the user owns the store associated with the product
	if product.Store.OwnerID != userID {
		return errors.New("unauthorized action: you do not own this store")
	}

	// Mark the product as deleted
	product.IsDeleted = true
	if err := db.GetDBConn().Save(&product).Error; err != nil {
		logger.Error.Printf("[repository.DeleteProduct] Error deleting product: %v\n", err)
		return err
	}

	return nil
}

// GetAllProducts retrieves all products filtered by category, price range, etc., and includes associated ProductImage data.
func GetAllProducts(minPrice, maxPrice float64, categoryID uint, productName string, storeID uint) ([]models.Product, error) {
	var products []models.Product

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

func CreateProductWithImages(product *models.Product, images []models.ProductImage, userID uint) error {
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
