package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup //Содаем вейт группу

	for _, v := range arr { // проходим по массиву
		wg.Add(1)        // добавляем в вейтгруппу горутину
		go func(v int) { //вызываем горутину
			fmt.Println(v * v) //выводим результат
			wg.Done()          // удаляем горутину из группы
		}(v)
	}
	wg.Wait() // ждем выполнения всех горутин
}
