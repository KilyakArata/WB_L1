package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	// Установка значений a и b
	a.SetString("123456789012345678901234567890", 10)
	b.SetString("987654321098765432109876543210", 10)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма: %s\n", sum.String())

	// Вычитание
	diff := new(big.Int).Sub(a, b)
	fmt.Printf("Разность: %s\n", diff.String())

	// Умножение
	prod := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение: %s\n", prod.String())

	// Деление
	div := new(big.Int).Div(a, b)
	fmt.Printf("Частное: %s\n", div.String())
}
