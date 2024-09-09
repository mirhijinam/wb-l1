/*
 * Имеется последовательность строк: (cat, cat, dog, cat, tree). 
 * Создать для нее собственное множество.
 */

package main

import "fmt"

func createSet(seq []string) []string {
	ret := make([]string, 0, len(seq))

	m := make(map[string]struct{})

	for _, x := range seq {
		m[x] = struct{}{}
	}

	for k, _ := range m {
		ret = append(ret, k)
	}

	return ret
}

func main() {
	fmt.Println(createSet([]string{"cat", "cat", "dog", "cat", "tree"}))
}
