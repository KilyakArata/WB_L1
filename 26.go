package main

import (
	"fmt"
	"strings"
)

func uniqueCharacters(s string) bool {
	charMap := make(map[rune]bool)

	// Приведем все символы к нижнему регистру для регистронезависимости
	s = strings.ToLower(s)

	for _, char := range s {
		if charMap[char] {
			return false // Если символ уже встречался, возвращаем false
		}
		charMap[char] = true
	}

	return true // Если все символы уникальны, возвращаем true
}

func main() {
	stringsToCheck := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, s := range stringsToCheck {
		fmt.Printf("Строка \"%s\": %v\n", s, uniqueCharacters(s))
	}
}
