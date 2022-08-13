package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	arg := os.Args[1:] // получаем количество секунд из аргументов
	if len(arg) < 1 {
		return
	}
	n, err := strconv.Atoi(arg[0])
	if err != nil || n < 1 {
		return
	}

	duration := time.Now().Add(time.Duration(n) * time.Second)
	// Создаем контекст с дедлайном. По истечении n секунд контекст отменится
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, duration)
	defer cancel()

	ch := make(chan int)
	defer close(ch)
	// Запускаем горутину которая выводит данные из канала пока кон
	go worker(ch, ctx)
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

func worker(ch chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // если контекст отменен завершаем горутину и выходим
			return
		case data := <-ch: // Выводим полученные в канал данные
			fmt.Println(data)
		}
	}
}
