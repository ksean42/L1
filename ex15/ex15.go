package main

import "fmt"

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

func createHugeString(len int) []int {
	res := make([]int, 0, len)
	for i := 0; i < 1024; i++ {
		res = append(res, i)
	}
	return res
}

func someFunc() {
	// При "перенарезке" создается новый слайс который ссылается на исходный.
	//Из-за этого огромный исходный слайс будет висеть в памяти и не может быть вычищен пока на него ссылается новый слайс
	v := createHugeString(1 << 10)
	justString := v[:100]
	fmt.Println(cap(justString))
	// Вместо этого лучше сделать копию нужного отрезка
	true := make([]int, 100)
	copy(true, v[:100])
	fmt.Println(cap(true))
}

func main() {
	someFunc()
}
