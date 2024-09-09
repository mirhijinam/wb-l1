package main

import "fmt"

func binSearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := (right + left) / 2
		if arr[mid] == target {
			return mid
		} else if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 3, 8, 13, 17, 22, 45, 47, 75}

	fmt.Println(binSearch(a, 11))
	fmt.Println(binSearch(a, 45))
}
