package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var number int64 = 2 // рандомное число
	arg := os.Args[1:]   // считываем из аргументов номер инвертируемого бита
	if len(arg) < 1 {
		return
	}
	n, err := strconv.Atoi(arg[0])
	if err != nil || n < 1 || n > 64 {
		fmt.Println("Please enter number from 1 to 64")
		return
	}
	invert := int64(1 << (n - 1)) // Создаем битовую маску. двигаем бит влево на n - 1 позиций
	number = number ^ invert      // Исключающее ИЛИ
	fmt.Println(number)
}
