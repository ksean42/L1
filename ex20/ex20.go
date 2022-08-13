package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun moon star"
	fmt.Println(reverse(str))
}

func reverse(str string) string {
	slice := strings.Fields(str)
	start := 0
	end := len(slice) - 1
	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}
	return strings.Join(slice, " ")
}
