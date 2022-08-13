package main

import "fmt"

func main() {
	arr := []int{5, 8, 10, 14, 16, 20, 24, 32} // Входные данные должны быть отсортированы.
	target := 20
	i := binarySearch(arr, 0, len(arr)-1, target)
	fmt.Println(i)
}

// Возвращает индекс искомого элемента. -1 если элемент не найден
func binarySearch(arr []int, low int, high int, target int) int {
	if high-low <= 0 && arr[high] != target { // Базовый случай. Если мы в конце ветки и искомый элемент не найден, возвращаем -1
		return -1
	}

	mid := (high + low) / 2 // Вычисляем середину
	if arr[mid] > target {  // Сравниваем искомый элемент со средним
		return binarySearch(arr, low, mid, target) // если средний элемент больше искомого - ищем слева
	} else if arr[mid] < target {
		return binarySearch(arr, mid+1, high, target) // если средний элемент меньше искомого - ищем справа
	} else { // иначе возвращаем средний
		return mid
	}
}
