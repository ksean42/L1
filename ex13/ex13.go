package main

import "fmt"

func main() {
	a := 5
	b := 23
	// Арифметический способ
	a = a + b // сумма двух чисел
	b = a - b // кладем в b разность суммы и b (28 - 23 = 5 = a)
	a = a - b // кладем в a разность суммы и нового b

	fmt.Printf("a = %d, b = %d\n", a, b)
}
