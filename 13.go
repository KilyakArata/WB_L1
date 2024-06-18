package main

import "fmt"

func main() {
	fmt.Println("Первый способ")
	a, b := 10, 20

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Обмен значениями
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

	fmt.Println("Второй способ")
	a, b = 10, 20

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	aSlice := []int{a}
	aSlice[0], b = b, aSlice[0]
	a = aSlice[0]

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
