package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter string:")
	fmt.Scan(&str)
	fmt.Println(isUnique(str))
}

func isUnique(str string) bool {
	m := make(map[rune]struct{})
	str = strings.ToLower(str) // Преобразуем к нижнему регистру
	r := []rune(str)           // преобразуем строку в слайс рун для поддержки юникода

	// Заносим каждый символ в мапу, если символ встретился более одного раза - возвращаем false
	for _, v := range r {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
