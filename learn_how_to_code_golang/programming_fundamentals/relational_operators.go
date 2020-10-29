package main

import "fmt"

func main() {

	var a, b = 42, 43
	c := a == b
	d := a >= b
	e := a <= b
	f := a != b
	g := a < b
	h := a > b
	fmt.Println(c, d, e, f, g, h)

}
