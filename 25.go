package main

import (
	"fmt"
	"time"
)

// Sleep реализует функцию ожидания на указанное количество миллисекунд.
func Sleep(milliseconds int) {
	<-time.After(time.Millisecond * time.Duration(milliseconds))
}

func main() {
	fmt.Println("Начало работы программы")
	Sleep(2000) // Ожидание 2 секунды
	fmt.Println("Прошло 2 секунды")
}
