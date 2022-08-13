package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	res := make(chan int)
	sum := make(chan int)
	var wg sync.WaitGroup

	go func(sumChan chan int, res chan int) { // запускаем в новой горутине функцию подсчета суммы квадратов
		sum := 0
		for {
			select {
			case i, ok := <-sumChan:
				if ok { // если канал еще не закрыт прибавляем
					sum += i
				} else {
					res <- sum // если канал закрылся передаем результат в канал результата
					return
				}
			}

		}
	}(sum, res)

	for _, v := range arr { // проходим по массиву
		wg.Add(1)        // добавляем в вейтгруппу горутину
		go func(v int) { //вызываем горутину
			sum <- v * v // передаем квадрат числа в канал
			wg.Done()    // удаляем горутину из группы
		}(v)
	}
	wg.Wait()
	close(sum)         //закрываем канал для отправки результата в канал res
	fmt.Println(<-res) // вывод результата
}
