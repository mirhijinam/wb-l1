/*
 * Разработать программу, которая будет последовательно отправлять значения 
 * в канал, а с другой стороны канала — читать.
 * По истечению N секунд программа должна завершаться.
 *
 * 3 goroutines: main, receiver, sender
 */

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sender(ctx context.Context, ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var x int
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- x
			time.Sleep(time.Millisecond * 200)
			x++
		}
	}
}

func receiver(ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	ch := make(chan interface{})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(n))
	defer func() {
		fmt.Println("program was stopped")
		cancel()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go sender(ctx, ch, wg)
	go receiver(ch, wg)

	wg.Wait()
}
