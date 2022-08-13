package main

import (
	"fmt"
	"math"
)

// Структура точки
type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Вычисляет евклидово расстояние
func distance(a, b *Point) float64 {
	return math.Sqrt(math.Pow((a.x-b.x), 2) + math.Pow((a.y-b.y), 2))
}

func main() {
	//Создаем две точки
	point := NewPoint(5, 10)
	point2 := NewPoint(2, 8)
	// Выводим результат выполнения функции подсчета расстояния
	fmt.Println(distance(point, point2))
}
