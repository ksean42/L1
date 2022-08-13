package main

import "fmt"

// Базовая структура
type Human struct {
	a int
}

type Action struct {
	Human // встраиваем методы и поля Human в другую структуру
}

// Метод структуры Human
func (h *Human) hello() {
	fmt.Printf("Hi from human struct! %d\n", h.a)
}

func main() {
	action := &Action{Human{1}} // Создаем экземпляры структур
	human := &Human{2}
	action.hello() // Вызываем методы
	human.hello()
}
