package main

import "fmt"

// Функция быстрой сортировки
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Выбираем опорный элемент (pivot)
	middle := arr[len(arr)/2]

	// Разделяем массив на три части
	left := []int{}
	right := []int{}
	equal := []int{}

	for _, val := range arr {
		if val < middle {
			left = append(left, val)
		} else if val > middle {
			right = append(right, val)
		} else {
			equal = append(equal, val)
		}
	}

	// Рекурсивно сортируем левую и правую части
	left = quickSort(left)
	right = quickSort(right)

	// Объединяем отсортированные части
	return append(append(left, equal...), right...)
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Неотсортированный массив:", arr)
	sortedArr := quickSort(arr)
	fmt.Println("Отсортированный массив:", sortedArr)

}
