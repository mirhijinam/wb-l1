/*
 * Реализовать пересечение двух неупорядоченных множеств.
 */

package main

import "fmt"

func intersect(a, b []int) []int {
	ret := make([]int, 0, min(len(a), len(b)))

	m := make(map[int]bool)

	for _, x := range a {
		m[x] = false
	}

	for _, y := range b {
		if v, ok := m[y]; ok && v == false {
			ret = append(ret, y)
			m[y] = true
		}
	}

	return ret
}

func main() {
	a := []int{1, 1, 1, 1}
	b := []int{1, 1, 1}
	fmt.Println(intersect(a, b))

	c := []int{1, 2, 3, 4, 4, 5, 5, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 15, 16}
	d := []int{12, 12, 12, 13, 14, 15, 16, 2, 19, 20, 18, 16, 17}
	fmt.Println(intersect(c, d))

}
