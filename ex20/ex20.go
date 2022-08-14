package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun moon star"
	fmt.Println(reverse(str))
	fmt.Println(reverse1(str))
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

func reverse1(str string) string {
	newStr := new(strings.Builder)
	newStr.Grow(len(str) + 1)
	slice := strings.Fields(str)
	for i := len(slice) - 1; i >= 0; i-- {
		newStr.WriteString(slice[i] + " ")
	}
	return newStr.String()
}
