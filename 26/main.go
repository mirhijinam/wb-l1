/*
 * Разработать программу, которая проверяет, что все символы в строке уникальные
 * (true — если уникальные, false etc). 
 * Функция проверки должна быть регистронезависимой.
 *
 * Например: 
 * abcd — true
 * abCdefAaf — false
 * aabcd — false
 */

package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	s = strings.ToLower(s) // Так как функция должна быть регистронезависимой

	uniqueMap := make(map[rune]struct{})
	for _, v := range []rune(s) {
		uniqueMap[v] = struct{}{}
	}
	if len([]rune(s)) != len(uniqueMap) {
		return false
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}
