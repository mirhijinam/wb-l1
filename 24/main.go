/*
 * Разработать программу нахождения расстояния между двумя точками,
 * которые представлены в виде структуры Point
 * с инкапсулированными параметрами x,y и конструктором.
 */

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}

func Distance(p1, p2 Point) float64 {
	distx := p1.GetX() - p2.GetX()
	disty := p1.GetY() - p2.GetY()

	return math.Sqrt(distx*distx + disty*disty)
}

func main() {
	p1 := New(1.123, 2.1231)
	p2 := New(5.567, 6.5675)
	fmt.Println(Distance(p1, p2))

	p3 := New(1, 1)
	p4 := New(4, 5)
	fmt.Println(Distance(p3, p4))
}
