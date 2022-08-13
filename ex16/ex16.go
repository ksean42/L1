package main

import "fmt"

func main() {
	arr := []int{9, 1, 3, 14, 7, 0, 30, 2, -1, 8}
	quicksort(arr)
	fmt.Println(arr)
}

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := swap(arr)
	fmt.Println(arr)
	quicksort(arr[:pivot])
	quicksort(arr[pivot+1:])
}

func swap(arr []int) int {
	high := len(arr) - 1
	pivot := arr[high]
	low := 0
	for i := range arr {
		if arr[i] < pivot {
			arr[low], arr[i] = arr[i], arr[low]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	return low
}
