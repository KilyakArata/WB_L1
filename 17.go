package main

import (
	"fmt"
)

// binarySearch выполняет бинарный поиск элемента в отсортированном массиве
func binarySearch(arr []int, target int) int {
	low := 0             // Начальный индекс
	high := len(arr) - 1 // Конечный индекс

	for low <= high {
		mid := (low + high) / 2 // Средний индекс
		if arr[mid] < target {
			low = mid + 1 // Ищем в правой половине
		} else if arr[mid] > target {
			high = mid - 1 // Ищем в левой половине
		} else {
			return mid // Элемент найден
		}
	}
	return -1 // Элемент не найден
}

func main() {
	// Пример использования бинарного поиска
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
