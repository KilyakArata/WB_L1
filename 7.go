package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем RWMutex для синхронизации доступа к карте
	var mu sync.RWMutex
	// Создаем карту для хранения данных
	data := make(map[int]int)
	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем 10 горутин для записи данных в карту
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины
			mu.Lock()       // Захватываем мьютекс для записи в карту
			data[i] = i * i // Записываем данные в карту
			mu.Unlock()     // Освобождаем мьютекс после записи
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Используем RLock для чтения данных из карты
	mu.RLock() // Захватываем мьютекс для чтения данных из карты
	for k, v := range data {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}
	mu.RUnlock() // Освобождаем мьютекс после чтения
}
