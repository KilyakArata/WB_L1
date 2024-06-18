package main

import (
	"fmt"
	"sync"
)

func main() {
	// Инициализируем массив чисел
	list := []int{2, 4, 6, 8, 10}

	// Создаем экземпляр sync.WaitGroup для синхронизации горутин
	var wg sync.WaitGroup

	// Переменная для хранения суммы квадратов
	sum := 0

	// Мьютекс для безопасного изменения переменной sum из нескольких горутин
	var mu sync.Mutex

	// Проходимся по массиву чисел
	for i := 0; i < len(list); i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup перед запуском каждой горутины

		// Запускаем горутину для расчета квадрата числа
		go func(i int) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup по завершении горутины

			// Рассчитываем квадрат числа
			square := list[i] * list[i]

			// Захватываем мьютекс перед изменением общей переменной sum
			mu.Lock()
			sum += square
			mu.Unlock() // Освобождаем мьютекс
		}(i) // Передаем текущее значение i в анонимную функцию, чтобы избежать условий гонки
	}

	wg.Wait() // Ожидаем завершения всех горутин

	// Выводим сумму квадратов чисел
	fmt.Println(sum)
}
