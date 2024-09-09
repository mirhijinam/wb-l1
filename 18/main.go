/*
 * Реализовать структуру-счетчик, 
 * которая будет инкрементироваться в конкурентной среде.
 * По завершению программа должна выводить итоговое значение счетчика.
 */

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    *sync.Mutex
}

func New() Counter {
	return Counter{
		value: 0,
		mu:    &sync.Mutex{},
	}
}

func (c *Counter) Increment() {
	c.value++
}

func main() {
	workers := 1337
	wg := &sync.WaitGroup{}

	c := New()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.mu.Lock()
			defer c.mu.Unlock()
			c.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(c.value)
}
