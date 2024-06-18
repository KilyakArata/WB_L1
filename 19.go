package main

import "fmt"

// Функция для переворота строки
func reverseString(input string) string {
	runes := []rune(input) // Преобразуем строку в срез рун
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i] // Меняем местами символы
	}
	return string(runes) // Преобразуем срез рун обратно в строку
}

func main() {
	var input string

	fmt.Println("Введите строку для переворота:")
	fmt.Scanln(&input) // Чтение входной строки

	reversed := reverseString(input) // Переворот строки

	fmt.Printf("Перевернутая строка: %s\n", reversed) // Вывод перевернутой строки
}
