/*
 * Разработать программу, которая перемножает, делит, складывает,
 * вычитает две числовых переменных a,b, значение которых > 2^20.
 */

package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := new(big.Int)
	x.SetString("221221221221221221221221221221", 10)

	y := new(big.Int)
	y.SetString("112112112112112112112112112112", 10)

	p := new(big.Int).Mul(x, y)
	fmt.Println(p.String())

	d := new(big.Rat).SetFrac(x, y)
	fmt.Println(d.FloatString(50))

	a := new(big.Int).Add(x, y)
	fmt.Println(a.String())

	s := new(big.Int).Sub(x, y)
	fmt.Println(s.String())
}
