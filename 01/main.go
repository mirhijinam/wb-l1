/*
 * Дана структура Human (с произвольным набором полей и методов). 
 * Реализовать встраивание методов в структуре Action 
 * от родительской структуры Human (аналог наследования).
 */

package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name string
	Age  int
}

func (h Human) Say(phrase string) {
	fmt.Println(phrase)
}

type Action struct {
	Human // Встраиваем методы Human в структуру Action
}

func (a Action) Shout(phrase string) {
	fmt.Println(strings.ToUpper(phrase))
}

func main() {
	a := &Action{}
	a.Say("test")
	a.Shout("test")
}
