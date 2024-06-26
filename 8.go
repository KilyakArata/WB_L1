package main

import "fmt"

func main() {
	var num int64
	var pos uint
	var bitValue int

	fmt.Println("Введите число:")
	fmt.Scan(&num) // Чтение значения переменной num

	fmt.Println("Введите позицию бита (от 0 до 63):")
	fmt.Scan(&pos) // Чтение позиции бита

	fmt.Println("Введите значение бита (0 или 1):")
	fmt.Scan(&bitValue) // Чтение значения бита

	// Печатаем число до изменения в двоичном формате
	fmt.Printf("Число до изменения:\n%064b\n", num)

	if bitValue == 1 {
		// Промежуточное число для установки бита в 1
		mask := int64(1 << pos)
		fmt.Printf("Маска для установки бита в 1:\n%064b\n", mask)
		// Устанавливаем pos-й бит в 1
		num = num | mask
	} else if bitValue == 0 {
		// Промежуточное число для установки бита в 0
		mask := int64(1 << pos)
		fmt.Printf("Маска для установки бита в 0:\n%064b\n", mask)
		// Устанавливаем pos-й бит в 0
		num = num &^ mask
	} else {
		fmt.Println("Неверное значение бита. Введите 0 или 1.")
		return
	}

	/*
		"<<" - Этот оператор сдвигает битовую единицу (1) влево на pos позиций.
		"|" - (битовый OR) сравнивает два числа побитно и устанавливает каждый
		бит результата в 1, если хотя бы один из соответствующих битов операндов равен 1.
		"&^" - оператор AND NOT , который устанавливает pos-й бит в 0, оставляя остальные биты неизменными.
	*/

	// Печатаем число после изменения в двоичном формате
	fmt.Printf("Число после изменения:\n%064b\n", num)
	fmt.Printf("Результат:\n%d\n", num) // Вывод результата
}
