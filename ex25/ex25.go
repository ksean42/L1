package main

import (
	"fmt"
	"time"
)

func main() {
	timeToSleep := 4

	// Способ1
	fmt.Println("Start 1....", time.Now())
	<-time.After(time.Second * time.Duration(timeToSleep))
	fmt.Println("Finish", time.Now())
	// Способ 2
	fmt.Println("Start 2....", time.Now())
	mySleep(time.Second * time.Duration(timeToSleep))
	fmt.Println("Finish", time.Now())

}

// В цикле сравниваем время
func mySleep(t time.Duration) {
	timeToSleep := time.Now().Add(t) // Время сна = Текущее время + продолжительность сна
	now := time.Now()
	for !now.After(timeToSleep) { // Не выходим из цикла пока текущее время не больше времени сна
		now = time.Now()
	}
}
