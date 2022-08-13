package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter string:")
	fmt.Scan(&str)
	r := []rune(str) // Ковертируем строку в руны для поддержки юникода
	start := 0
	end := len(r) - 1
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
	fmt.Println(string(r))
}

// главрыба且丕
