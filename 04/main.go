/*
 * Реализовать постоянную запись данных в канал (главный поток).
 * Реализовать набор из N воркеров, 
 * которые читают произвольные данные из канала и выводят в stdout.
 * Необходима возможность выбора количества воркеров при старте.
 *
 * Программа должна завершаться по нажатию Ctrl+C. 
 * Выбрать и обосновать способ завершения работы всех воркеров.
 */

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var n int
	fmt.Scanf("%d", &n)

	// Определение канала с произвольным типом данных
	ch := make(chan interface{})

	wg := &sync.WaitGroup{}

	// Создание пула горутин-ридеров
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				fmt.Printf("id=%d,msg=%v\n", i, v) // Чтение из канала
			}
		}()
		// Имеется доступ к каналу через замыкание, т.к. находятся с ним в одной области видимости
		// В новых версиях языка необязательно захватывать переменную-счетчик
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigch
		defer cancel()
		fmt.Println("\nshutting down...")
		close(ch)
	}()

	// Горутина, записывающая данные в канал
	go func() {
		var x int
		for {
			select {
			case <-ctx.Done():
				break
			default:
				ch <- x
				// Имитация полезной работы, не имеет отношения к синхронизации
				time.Sleep(time.Millisecond * 200)
				x++
			}
		}
	}()

	wg.Wait()
	fmt.Println("work was done")
}
