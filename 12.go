package main

import "fmt"

func main() {
	// Данная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание множества из строк
	set := make(map[string]struct{})
	for _, v := range strings {
		set[v] = struct{}{}
	}

	// Вывод множества
	fmt.Println("Множество строк:")
	for k := range set {
		fmt.Println(k)
	}
}
