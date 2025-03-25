package db

import (
	"BizMart/internal/app/models"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	ctx         = context.Background()
)

// InitializeRedis инициализирует соединение с Redis
func InitializeRedis(redisParams models.RedisParams) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisParams.Host, redisParams.Port), // адрес Redis-сервера
		Password: redisParams.Password,                                     // если пароль не установлен, оставьте пустым
		DB:       redisParams.DB,                                           // используемая база данных Redis
	})

	// Проверка соединения
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
		return err
	}

	return nil
}

// SetCache записывает данные в кэш с указанным сроком жизни
func SetCache(key string, value interface{}, expiration time.Duration) error {
	err := RedisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Printf("Error setting cache in Redis: %v", err)
		return err
	}
	return nil
}

// GetCache получает данные из кэша по ключу
func GetCache(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Printf("Key does not exist in Redis: %s", key)
		return "", nil
	} else if err != nil {
		log.Printf("Error getting cache from Redis: %v", err)
		return "", err
	}
	return val, nil
}

// DeleteCache удаляет данные из кэша по ключу
func DeleteCache(key string) error {
	err := RedisClient.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Error deleting cache from Redis: %v", err)
		return err
	}
	return nil
}

func CloseRedisConnection() error {
	err := RedisClient.Close()
	if err != nil {
		return err
	}

	return nil
}
