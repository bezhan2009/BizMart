package jobs

import (
	"encoding/json"
	"log"
	"time"

	models2 "BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/db"
)

const cacheKey = "cachedProducts"

// UpdateProductCache обновляет кэш продуктов каждые 10 минут
func UpdateProductCache() {
	update := func() {
		products, err := repository.GetAllProducts(0, 0, 0, "", 0) // Параметры фильтрации по умолчанию
		if err != nil {
			log.Printf("Error updating product cache: %v", err)
			return
		}

		// Сериализация продуктов в JSON
		productData, err := json.Marshal(products)
		if err != nil {
			log.Printf("Error marshaling products: %v", err)
			return
		}

		// Запись данных в Redis
		err = db.SetCache(cacheKey, productData, 10*time.Minute)
		if err != nil {
			log.Printf("Error setting cache in Redis: %v", err)
			return
		}
	}

	update()

	ticker := time.NewTicker(60 * time.Minute)
	for {
		select {
		case <-ticker.C:
			update()
		}
	}
}

// GetCachedProducts возвращает кэшированные продукты
func GetCachedProducts() ([]models2.Product, error) {
	// Получение данных из Redis
	productData, err := db.GetCache(cacheKey)
	if err != nil {
		return nil, err
	}
	if productData == "" {
		return nil, nil
	}

	// Десериализация данных из JSON в структуру продуктов
	var products []models2.Product
	err = json.Unmarshal([]byte(productData), &products)
	if err != nil {
		log.Printf("Error unmarshaling products: %v", err)
		return nil, err
	}

	return products, nil
}
