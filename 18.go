package main

import (
	"fmt"
	"sync"
)

// Counter - структура-счетчик
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment - метод для инкрементации счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value - метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// Количество горутин
	numGoroutines := 100

	// Запуск множества горутин для инкрементации счетчика
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод итогового значения счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}
