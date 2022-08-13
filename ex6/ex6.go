package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// example1()
	// example2()
	example3()
}

func example1() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	go func(ctx context.Context, wg *sync.WaitGroup) {
		n := 0
		wg.Add(1)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancel")
				wg.Done()
				return
			default:
				fmt.Println(n)
				time.Sleep(500 * time.Millisecond)
				n++
			}
		}

	}(ctx, &wg)
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}

func example2() {
	duration := time.Now().Add(2 * time.Second)
	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, duration)

	wg := sync.WaitGroup{}
	ch := make(chan int, 1)
	go func(ctx context.Context, ch chan int) {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- n
				time.Sleep(500 * time.Millisecond)
				n++
			}
		}
	}(ctx, ch)
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancel with deadline")
				wg.Done()
				return
			case n := <-ch:
				fmt.Println(n)
			}
		}

	}(ctx, &wg, ch)
	wg.Wait()
}

func example3() {
	ch := make(chan string)
	exit := make(chan bool)
	go func(ch chan string, exit chan bool) {
		for {
			select {
			case <-exit:
				close(ch)
				return
			default:
				ch <- "Hello!"
				time.Sleep(time.Second)
			}
		}
	}(ch, exit)
	go func(ch chan string) {
		for {
			if m, open := <-ch; open {
				fmt.Println(m)
			} else {
				fmt.Println("Stop by closing channel")
				return
			}
		}
	}(ch)
	time.Sleep(5 * time.Second)
	exit <- true
}
