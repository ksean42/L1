package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 5
	var b interface{} = "string"
	var c interface{} = true
	var d interface{} = make(chan int)

	fmt.Println(getType(a))
	fmt.Println(getType(b))
	fmt.Println(getType(c))
	fmt.Println(getType(d))
}

func getType(input interface{}) string {
	/*
		С помощью switch
	*/
	// switch input.(type) {
	// 	case int:
	// 		return "int"
	// 	case string:
	// 		return "string"
	// 	case bool:
	// 		return "bool"
	// 	case chan int:
	// 		return "chan int"
	// 	default:
	// 		return "another type"
	// }

	/*
		С помощью sprintf
	*/

	// return (fmt.Sprintf("%T", input))

	/* С помощью рефлексии
	Рефлексия — способность программы исследовать собственную структуру, в особенности через типы
	*/

	return reflect.TypeOf(input).String()
}
