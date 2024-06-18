package main

import "fmt"

// Определяем структуру Human с полями HumanName, Surname и Age
type Human struct {
	HumanName string
	Surname   string
	Age       int
}

// Определяем структуру Action с полем ActionName и встраиваем в нее структуру Human
type Action struct {
	ActionName string
	Human
}

// Метод CallName для структуры Human, который возвращает имя человека
func (h *Human) CallName() string {
	return fmt.Sprint(h.HumanName)
}

// Метод CallAge для структуры Human, который возвращает возраст человека
func (h *Human) CallAge() string {
	return fmt.Sprint(h.Age)
}

// Метод CallAction для структуры Action, который возвращает название действия
func (a *Action) CallAction() string {
	return fmt.Sprint(a.ActionName)
}

func main() {
	// Создаем экземпляр структуры Human с заданными значениями полей
	human := Human{
		HumanName: "Ann",
		Surname:   "Ivanova",
		Age:       60,
	}

	// Создаем экземпляр структуры Action с заданным значением поля ActionName и встраиваем в него экземпляр Human
	action := Action{
		ActionName: "Reading",
		Human:      human,
	}

	// Вызываем метод CallAction для экземпляра Action и сохраняем результат в переменную a
	a := action.CallAction()

	// Вызываем метод CallName для экземпляра Action, который доступен благодаря встраиванию Human, и сохраняем результат в переменную b
	b := action.CallName()

	// Вызываем метод CallAge для экземпляра Action, который доступен благодаря встраиванию Human, и сохраняем результат в переменную c
	c := action.CallAge()

	// Выводим значения переменных a, b и c
	fmt.Println(a, b, c)
}
