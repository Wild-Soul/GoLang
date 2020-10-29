package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func info(sh shape) {
	fmt.Println("Area of the given shape", sh.area())
}

func main() {
	s := square{
		side: 4,
	}
	c := circle{
		radius: 3,
	}
	info(s)
	info(c)
}
