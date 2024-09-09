/*
 * Дана последовательность чисел: 2,4,6,8,10. 
 * Найти сумму их квадратов: 2^2 + 4^2 + 6^2 + 8^2 + 10^2
 * с использованием конкурентных вычислений.
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	fn := func(wg *sync.WaitGroup, mu *sync.Mutex, x int) {
		defer wg.Done()
		mu.Lock() // Опускаем мьютекс для предотвращения состояния гонки
		count += x * x
		mu.Unlock() // Поднимаем мьютекс для предоставления доступа следующей горутине
	}

	arr := [5]int{2, 4, 6, 8, 10}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{} // Инициализируем инструмент синхронизации

	for _, v := range arr {
		wg.Add(1)
		go fn(wg, mu, v)
	}

	wg.Wait()
	fmt.Println(count)
}
