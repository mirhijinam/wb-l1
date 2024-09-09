/*
 * Разработать программу, которая переворачивает подаваемую на ход строку 
 * (например: «главрыба — абырвалг»).
 * Символы могут быть unicode.
 */

package main

import (
	"fmt"
)

func flip(s string) string {
	res := make([]rune, len(s))
	sArr := []rune(s)

	for i := len(sArr) - 1; i >= 0; i-- {
		res[len(sArr)-i] = sArr[i]
	}

	return string(res)
}

func main() {
	str := "главрыба"
	fmt.Printf("%s - %s\n", str, flip(str))
}
