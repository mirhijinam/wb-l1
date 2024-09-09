/*
 * Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
 */

package main

import "fmt"

func quickSort(arr *[]int, low, high int) {
	if low < high {
		p := partition(*arr, low, high)
		fmt.Printf("arr[%v]=%v\n", p, (*arr)[p])
		fmt.Printf("%v\n", *arr)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	x := []int{12, -13, 14, 13, 88, -19, 37}
	quickSort(&x, 0, len(x)-1)
}
