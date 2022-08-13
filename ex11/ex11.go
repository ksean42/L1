package main

import "fmt"

func main() {
	a := []int{1, 5, 3, 9, 11, 54}
	b := []int{4, 1, 6, 0, 41, 9, 42, 11, 65}
	fmt.Println(firstSolution(a, b))
	secondSolution(a, b)
}

func firstSolution(a, b []int) map[int]bool {
	intersec := make(map[int]bool, len(a)) // инициализируем пустую карту

	for _, v := range a {
		intersec[v] = false //заносим ключи из левого множества
	}

	for _, v := range b {
		if _, ok := intersec[v]; ok {
			intersec[v] = true //при пересечении устанавливаем соответствующий флаг
		}
	}
	for k, v := range intersec {
		if v {
			fmt.Println(k) //выводим пересечения
		}
	}
	return intersec
}

// Предпочтительное решение
func secondSolution(a, b []int) []int {
	intersec := make(map[int]struct{})
	result := []int{} // слайс с пересечениями

	for _, v := range a {
		intersec[v] = struct{}{} // заносим ключи из левого множества, в качестве значения - пустая структура размером 0
	}

	for _, v := range b { // проходим по правому множеству
		if _, ok := intersec[v]; ok { // если элемент правого множества имеется в карте - добавляем в слайс пересечений
			result = append(result, v)
		}
	}

	fmt.Println(result)
	return result
}
