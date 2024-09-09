/*
 * Поменять местами два числа без создания временной переменной.
 */

package main

import "fmt"

func main() {
	a, b := 13, 37
	fmt.Println(a, b)
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)
}
