/*
 * Написать программу, которая конкурентно рассчитает значение квадратов чисел,
 * взятых из массива (2,4,6,8,10), и выведет их квадраты в stdout.
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	fn := func(wg *sync.WaitGroup, x int) {
		defer wg.Done() // Уведомляем о завершении горутины
		fmt.Println(x * x)
	}

	arr := [5]int{2, 4, 6, 8, 10}

	wg := &sync.WaitGroup{} // Инициализируем инструмент синхронизации

	for _, v := range arr {
		wg.Add(1) // Увеличиваем количество учитываемых горутин
		go fn(wg, v)
	}

	wg.Wait() // Ожидаем завершения всех горутин перед завершение основной горутины
}
