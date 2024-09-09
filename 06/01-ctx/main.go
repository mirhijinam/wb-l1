/*
 * Реализовать все возможные способы остановки выполнения горутины.
 *
 * Метод: пакет context
 */

package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("i am not working.")
			return
		default:
			fmt.Println("i am working!")
			time.Sleep(time.Millisecond * 333)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go foo(ctx)

	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Millisecond * 500)

}
