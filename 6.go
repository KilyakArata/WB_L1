package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Первый способ: использование канала для остановки горутины
	fmt.Println("Первый способ: использование канала для остановки горутины")
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				// Получаем сигнал для остановки через канал stop
				fmt.Println("Горутина под номером 1 остановлена")
				return
			default:
				// Выполнение работы горутиной
				fmt.Println("1 Горутина работает")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Ждем 2 секунды перед остановкой горутины
	time.Sleep(2 * time.Second)
	// Отправляем сигнал для остановки горутины через канал stop
	stop <- struct{}{}
	// Ждем 1 секунду перед завершением программы
	time.Sleep(1 * time.Second)
	// Закрываем канал stop
	close(stop)

	// Второй способ: использование контекста для остановки горутины
	fmt.Println("Второй способ: использование контекста для остановки горутины")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// Получаем сигнал для остановки через контекст
				fmt.Println("Горутина под номером 2 остановлена")
				return
			default:
				// Выполнение работы горутиной
				fmt.Println("2 Горутина работает")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	// Ждем 2 секунды перед остановкой горутины
	time.Sleep(2 * time.Second)
	// Отправляем сигнал для остановки горутины через контекст
	cancel()
	// Ждем 1 секунду перед завершением программы
	time.Sleep(1 * time.Second)

	// Третий способ: использование таймера для остановки горутины
	fmt.Println("Третий способ: использование таймера для остановки горутины")
	timer := time.NewTimer(2 * time.Second)

	go func() {
		for {
			select {
			case <-timer.C:
				// Получаем сигнал для остановки через таймер
				fmt.Println("Горутина под номером 3 остановлена")
				return
			default:
				// Выполнение работы горутиной
				fmt.Println("3 Горутина работает")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Ждем 3 секунды перед завершением программы
	time.Sleep(3 * time.Second)

	// Четвертый способ: использование канала для отправки данных и закрытия канала для остановки
	fmt.Println("Четвертый способ: использование канала для отправки данных и закрытия канала для остановки")
	ch := make(chan string)

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				// Получаем сигнал для остановки через закрытие канала ch
				fmt.Println("Горутина под номером 4 остановлена")
				return
			}
			// Выполнение работы горутиной
			fmt.Println(v)
		}
	}()

	// Отправляем данные в канал ch
	ch <- "4 Горутина работает"
	time.Sleep(500 * time.Millisecond)
	ch <- "4 Горутина работает"
	time.Sleep(500 * time.Millisecond)
	// Закрываем канал ch
	close(ch)
	// Ждем 1 секунду перед завершением программы
	time.Sleep(time.Second)
}
