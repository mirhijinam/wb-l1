/*
 * Удалить i-ый элемент из слайса.
 */

package main

import (
	"fmt"
	"slices"
)

func main() {
	i := 1

	sl := []int{1, 2, 3, 4, 5, 6}
	sl = slices.Delete(sl, i, i+1)
	fmt.Println(sl)

	sl = []int{1, 2, 3, 4, 5, 6}
	sl = append(sl[:i], sl[i+1:]...)
	fmt.Println(sl)

}
