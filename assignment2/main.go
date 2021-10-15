package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func main() {
	ex1 := triangle{base: 10.0, height: 5.0}
	ex2 := square{sideLength: 40.4}

	printArea(ex1) // Output == 25.0
	printArea(ex2) // Output == 1632.16
}