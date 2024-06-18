package main

import "fmt"

func main() {
	// Инициализация переменных различных типов
	a := 1              // int
	b := "string"       // string
	c := true           // bool
	d := make(chan int) // channel of int
	f := struct{}{}     // пустая структура

	// Вызов функции check для каждой переменной
	fmt.Println("Первый способ")
	check(a)
	check(b)
	check(c)
	check(d)
	check(f)

	fmt.Println("Второй способ")
	Print(a)
	Print(b)
	Print(c)
	Print(d)
	Print(f)
}

// Определяем наш собственный constraint.
type Example interface {
	int | string | chan int | bool | struct{}
}

// Generic функция, которая работает с любым типом.
func Print[T Example](s T) {
	fmt.Printf("Переменная х имеет тип %T, значение: %v\n", s, s)
}

// Функция check определяет тип переменной и выводит соответствующее сообщение
func check(x interface{}) {
	// Использование type switch для определения типа переменной x
	switch v := x.(type) {
	case int:
		// Если тип переменной int
		fmt.Printf("Переменная x имеет тип int, значение: %d\n", v)
	case string:
		// Если тип переменной string
		fmt.Printf("Переменная x имеет тип string, значение: %s\n", v)
	case bool:
		// Если тип переменной bool
		fmt.Printf("Переменная x имеет тип bool, значение: %t\n", v)
	case chan int:
		// Если тип переменной channel of int
		fmt.Printf("Переменная x имеет тип chan int\n")
	default:
		// Если тип переменной не соответствует ни одному из перечисленных
		fmt.Printf("Переменная x имеет неизвестный тип\n")
	}
}
