package main

import (
	"fmt"
	"strconv"
	"sync"
)

type concurrentMap struct {
	sync.RWMutex // встраиваем мьютекс в структуру
	data         map[string]string
}

// Метод для записи в мапу
func (c *concurrentMap) write(str string) {
	c.Lock() // Блокируем на запись. Другие горутины будут ждать разблокировки прежде чем начать запись
	c.data[str] = str
	c.Unlock()
}

func main() {
	m := &concurrentMap{data: make(map[string]string)}
	wg := &sync.WaitGroup{}

	// Конкурентно записываем данные в мапу
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			m.write(strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
	i := 1
	for k, v := range m.data {
		fmt.Printf("%s = %s   -%d\n", k, v, i)
		i++
	}
}
