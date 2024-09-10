package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(input string) string {
	hash := sha256.New()                   // Создаем новый SHA-256 хеш
	hash.Write([]byte(input))              // Добавляем строку в хеш
	hashedBytes := hash.Sum(nil)           // Получаем итоговый хеш в виде байтового среза
	return hex.EncodeToString(hashedBytes) // Конвертируем байты в строку в формате hex
}
