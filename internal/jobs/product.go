package jobs

import (
	models2 "BizMart/internal/app/models"
	"BizMart/internal/repository"
	"log"
	"sync"
	"time"
)

var (
	cachedProducts []models2.Product
	mu             sync.RWMutex
)

// UpdateProductCache обновляет кэш продуктов каждые 10 минут
func UpdateProductCache() {
	// Функция для обновления данных
	update := func() {
		products, err := repository.GetAllProducts(0, 0, 0, "", 0) // Параметры фильтрации по умолчанию
		if err != nil {
			log.Printf("Error updating product cache: %v", err)
			return
		}

		// Обновляем кэш с блокировкой
		mu.Lock()
		cachedProducts = products
		mu.Unlock()
	}

	// Обновляем сразу при запуске
	update()

	// Запускаем обновление каждые 10 минут
	ticker := time.NewTicker(10 * time.Minute)
	for {
		select {
		case <-ticker.C:
			update()
		}
	}
}

// GetCachedProducts возвращает кэшированные продукты
func GetCachedProducts() []models2.Product {
	mu.RLock()
	defer mu.RUnlock()
	return cachedProducts
}
