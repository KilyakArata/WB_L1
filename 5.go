package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Запрашиваем у пользователя количество секунд для работы программы
	var duration int
	fmt.Println("Введите количество секунд для работы программы:")
	_, err := fmt.Scan(&duration)
	if err != nil {
		panic(err) // Завершаем программу в случае ошибки ввода
	}

	// Создаем канал для передачи данных
	a := make(chan int)

	// Создаем контекст с таймаутом для завершения работы через N секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel() // Обеспечиваем вызов cancel() при завершении функции main для освобождения ресурсов

	// Горутина для записи данных в канал
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done(): // Если контекст завершен (по истечению времени)
				close(a) // Закрываем канал, чтобы завершить работу горутины чтения
				return
			default:
				i++
				a <- i                             // Отправляем данные в канал
				time.Sleep(500 * time.Millisecond) // Задержка для имитации работы
			}
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		for v := range a {
			// Читаем данные из канала и выводим их на экран
			fmt.Printf("Получено значение: %d\n", v)
		}
	}()

	// Ожидание завершения контекста
	<-ctx.Done()
	fmt.Println("Программа завершена")
}
