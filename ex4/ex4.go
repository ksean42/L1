package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	arg := os.Args[1:] // Получаем количество воркеров из аргументов
	if len(arg) < 1 {
		return
	}
	n, err := strconv.Atoi(arg[0])
	if err != nil || n < 1 {
		return
	}
	// Создаем контекст
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan int)
	defer close(ch)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)

	go func() { // Запускаем горутину которая перехватывает сигнал выхода и отменяет контекст
		<-exit
		cancel()
	}()
	// Запускаем n воркеров
	for i := 1; i <= n; i++ {
		go worker(ch, i, ctx)
	}
	// Публикуем данные в канал пока контекст не отменен
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Terminated")
			return
		default:
			ch <- rand.Int()
			time.Sleep(time.Second)
		}

	}
}

func worker(ch chan int, n int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // если контекст отменен завершаем горутину и выходим
			fmt.Printf("Worker number %d shouting down\n", n)
			return
		case data := <-ch: // Выводим полученные в канал данные
			fmt.Println(data)
		}
	}
}
