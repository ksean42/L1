package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	var str string
	//Считываем числа со стандартного ввода
	fmt.Println("Enter number a:")
	fmt.Scan(&str)
	a, ok := new(big.Int).SetString(str, 10)
	if !ok {
		log.Fatal("Error.")
	}
	fmt.Println("Enter number b:")
	fmt.Scan(&str)
	b, ok := new(big.Int).SetString(str, 10)
	if !ok {
		log.Fatal("Error.")
	}
	// Производим операции с числами
	fmt.Print("A * B = ")
	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Print("A / B = ")
	fmt.Println(new(big.Int).Div(a, b))
	fmt.Print("A + B = ")
	fmt.Println(new(big.Int).Add(a, b))
	fmt.Print("A - B = ")
	fmt.Println(new(big.Int).Sub(a, b))
}
