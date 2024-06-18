package main

import (
	"fmt"
	"unicode/utf8"
)

/*
1 проблема:
createHugeString создает очень большую строку, то вся эта строка будет оставаться в памяти, потому что срез v[:100]
все еще ссылается на исходный массив байтов. Это приведет к утечке памяти.

2 проблема
Если строка v содержит символы, представленные несколькими байтами (например, символы Unicode), срез v[:100]
может нарушить границы символов и привести к некорректному отображению строки.
*/

var justString string

// createHugeString создает строку заданной длины
func createHugeString(size int) string {
	// Здесь можно генерировать строку любой длины. Для примера создадим строку из букв "a".
	return string(make([]rune, size))
}

func someFunc() {
	v := createHugeString(1 << 10)
	// Сначала создаем срез из 100 байт, но затем учитываем, что символы могут быть больше одного байта
	if utf8.RuneCountInString(v) < 100 {
		justString = v
	} else {
		// Создаем срез на первых 100 символов
		justString = string([]rune(v)[:100])
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
