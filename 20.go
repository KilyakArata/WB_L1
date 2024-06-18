package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция для переворота слов в строке
func reverseWords(input string) string {
	words := strings.Fields(input) // Разбиваем строку на слова
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i] // Меняем местами слова
	}
	return strings.Join(words, " ") // Объединяем слова обратно в строку
}

func main() {
	var input string

	fmt.Println("Введите строку для переворота слов:")
	// Чтение входной строки. Используем Scanln для ввода с пробелами
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	reversed := reverseWords(input) // Переворот слов в строке

	fmt.Printf("Строка с перевернутыми словами: %s\n", reversed) // Вывод результата
}
