/*
 * Разработать программу, которая переворачивает слова в строке.
 * Пример: «snow dog sun — sun dog snow».
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	arr := strings.Split(s, " ")

	result := make([]string, len(arr))
	for i, v := range arr {
		result[len(result)-i-1] = v
	}
	return strings.Join(result, " ")
}

func main() {
	inpString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic("failed to read string")
	}

	inpString = inpString[:len(inpString)-1]

	fmt.Printf("%s - %s\n", inpString, reverseString(inpString))
}
