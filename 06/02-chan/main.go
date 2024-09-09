/*
 * Реализовать все возможные способы остановки выполнения горутины.
 *
 * Метод: chan
 */

package main

import (
	"fmt"
	"time"
)

func bar(stopch chan struct{}) {
	for {
		select {
		case <-stopch:
			fmt.Println("i am not working.")
			return
		default:
			time.Sleep(time.Millisecond * 333)
			fmt.Println("i am working!")
		}
	}
}

func main() {
	stopch := make(chan struct{})

	go bar(stopch)

	time.Sleep(time.Second * 3)
	stopch <- struct{}{}
}
