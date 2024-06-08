package randnum

import (
	"math/rand"
	"time"
)

// Функция для генерации 6-ти значного случайного числа
func GenerateRandomNumber() int {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Генерируем случайное число в диапазоне от 100000 до 999999 (включительно)
	return rand.Intn(900000) + 100000
}
