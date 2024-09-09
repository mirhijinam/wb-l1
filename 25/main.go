/*
 * Реализовать собственную функцию sleep.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	n := 3
	fmt.Println("hi")
	sleep(time.Second * time.Duration(n))
	fmt.Println("bye")
}

func sleep(d time.Duration) {
	<-time.After(d)
}
