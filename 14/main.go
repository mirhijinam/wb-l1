/*
 * Разработать программу, которая в рантайме способна определить
 * тип переменной: int, string, bool, channel из переменной типа interface{}.
 */

package main

import (
	"fmt"
	"reflect"
)

func typeOf(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func main() {
	var x interface{}

	x = 3
	fmt.Println(typeOf(x))
	x = "x"
	fmt.Println(typeOf(x))
	x = make(chan int)
	fmt.Println(typeOf(x))

}
