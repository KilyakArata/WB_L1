package main

import "fmt"

func main() {
	// Два неупорядоченных множества
	set1 := []int{2, 4, 3, 5, 1}
	set2 := []int{3, 6, 4, 5, 7}

	// Преобразование слайсов в множества
	map1 := make(map[int]struct{})
	for _, v := range set1 {
		map1[v] = struct{}{}
	}
	map2 := make(map[int]struct{})
	for _, v := range set2 {
		map2[v] = struct{}{}
	}

	// Нахождение пересечения множеств
	var result []int
	for k := range map1 {
		if _, found := map2[k]; found {
			result = append(result, k)
		}
	}

	// Вывод результата
	fmt.Println("Пересечение множеств:", result)
}
