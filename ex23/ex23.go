package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 3
	fmt.Println(arr, len(arr), cap(arr))

	//С помощью append
	arr = append(arr[:i], arr[i+1:]...) // Сдвигаем  все элементы после i-го влево на один элемент
	fmt.Println(arr, len(arr), cap(arr))

	//В цикле
	arr2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 = delByIndex(arr2, i)
	fmt.Println(arr2, len(arr2), cap(arr2))

	//Удаление без сохранения порядка элементов
	arr3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr3 = del(arr3, i)
	fmt.Println(arr3, len(arr3), cap(arr3))
}

func delByIndex(arr []int, i int) []int {
	result := make([]int, 0, len(arr)-1)

	for n, v := range arr {
		if n != i {
			result = append(result, v) // Помещаем в новый слайс все элементы кроме i-го
		}
	}
	return result
}

func del(arr []int, n int) []int {
	len := len(arr)
	arr[len-1], arr[n] = arr[n], arr[len-1] // Меняем местами удаляемый элемент с последним
	return arr[:len-1]                      // Возвращаем срез до последнего элемента
}
