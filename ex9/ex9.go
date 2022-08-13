package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10, 12}
	x := make(chan int, len(arr))
	square := make(chan int, len(arr))
	wg := &sync.WaitGroup{}

	for _, v := range arr {
		x <- v
	}
	close(x)
	wg.Add(2)
	go publisher(x, square, wg)
	go reciever(square, wg)
	wg.Wait()
}

func publisher(x chan int, square chan int, wg *sync.WaitGroup) {
	for {
		if i, ok := <-x; ok {
			square <- i * i
		} else {
			wg.Done()
			close(square)
			return
		}
	}
}

func reciever(square chan int, wg *sync.WaitGroup) {
	for {
		if i, ok := <-square; ok {
			fmt.Println(i)
		} else {
			wg.Done()
			return
		}
	}
}
