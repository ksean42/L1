package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{}) // инициализация пустой мапы

	for _, v := range arr { // итерируемся по последовательности
		if _, ok := set[v]; !ok { // если такой строки в мапе нет - инициализируем ключ с пустой структурой в качестве значения
			set[v] = struct{}{}
		}
	}

	for k, _ := range set {
		fmt.Printf("k = %s\n", k) // выводим множество
	}
}
