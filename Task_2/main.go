package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	tr := triangle{height: 10, base: 5}
	sq := square{sideLength: 5}

	printArea(tr)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.height * t.base) / 2
}

func (s square) getArea() float64 {
	return math.Pow(2, s.sideLength)
}
