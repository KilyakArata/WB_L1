package main

import "fmt"

func main() {
	// Массив чисел для обработки
	numbers := []int{1, 2, 3, 4, 5}

	// Создание двух каналов
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Запуск горутины для записи чисел в первый канал
	go func(numbers []int, ch1 chan int) {
		// Проход по массиву чисел
		for _, num := range numbers {
			ch1 <- num // Отправка числа в канал ch1
		}
		close(ch1) // Закрытие канала ch1 после записи всех чисел
	}(numbers, ch1)

	// Запуск горутины для умножения чисел и записи во второй канал
	go func(ch1 chan int, ch2 chan int) {
		// Чтение чисел из канала ch1
		for num := range ch1 {
			ch2 <- num * 2 // Умножение числа на 2 и отправка в канал ch2
		}
		close(ch2) // Закрытие канала ch2 после завершения чтения из ch1
	}(ch1, ch2)

	// Чтение чисел из второго канала и вывод их в stdout
	for num := range ch2 {
		fmt.Println(num) // Вывод чисел из канала ch2
	}
}
