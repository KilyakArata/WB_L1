package main

import "fmt"

func main() {
	// Исходная последовательность температурных колебаний
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создание карты для хранения групп температур с шагом 10 градусов
	groups := make(map[int][]float64)

	// Объединение значений в группы
	for _, temp := range temperatures {
		key := int(temp/10) * 10 // Определение ключа для группы
		groups[key] = append(groups[key], temp)
	}

	// Вывод группированных значений
	for key, temps := range groups {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
