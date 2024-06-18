package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2 // Индекс элемента, который нужно удалить

	// Удаление элемента из слайса
	if i < len(slice) {
		slice = append(slice[:i], slice[i+1:]...)
	} else {
		fmt.Println("Индекс выходит за границы слайса")
	}

	// Вывод слайса после удаления элемента
	fmt.Println(slice)
}
