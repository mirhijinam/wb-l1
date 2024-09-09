/*
 * Разработать конвейер чисел.
 * Даны два канала: в первый пишутся числа (x) из массива,
 * во второй — результат операции x*2.
 * После чего данные из второго канала должны выводиться в stdout.
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	firstch := make(chan int)
	lastch := make(chan int)

	wg.Add(2)

	// Промежуточная между двумя каналами горутинa
	go func(fch <-chan int, lch chan<- int) {
		defer wg.Done()
		for x := range fch {
			lch <- x * 2
		}
		close(lch) // Закрываем канал после завершения
	}(firstch, lastch)

	// Горутинa для вывода результата
	go func(outch <-chan int) {
		defer wg.Done()
		for v := range outch {
			fmt.Println(v)
		}
	}(lastch)

	// Отправляем данные в первый канал
	for _, v := range arr {
		firstch <- v
	}
	close(firstch) // Закрываем канал после отправки всех данных

	wg.Wait()
}
