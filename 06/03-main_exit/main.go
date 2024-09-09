/*
 * Реализовать все возможные способы остановки выполнения горутины.
 *
 * Метод: main exit
 */
package main

import (
	"fmt"
	"os"
	"time"
)

func foo() {
	for {
		time.Sleep(time.Millisecond * 333)
		fmt.Println("i am working!")
	}
}

func bar() {
	for {
		time.Sleep(time.Millisecond * 333)
		fmt.Println("i am working too!")
	}
}

func main() {
	go foo()
	go bar()

	time.Sleep(time.Second * 3)
	os.Exit(0)

}
