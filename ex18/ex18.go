package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	c int
}

func (c *Counter) increment() {
	c.Lock()
	defer c.Unlock()
	c.c++
}

func main() {
	counter := new(Counter)

	wg := sync.WaitGroup{}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(c *Counter, wg *sync.WaitGroup) {
			c.increment()
			wg.Done()
		}(counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.c)
}
