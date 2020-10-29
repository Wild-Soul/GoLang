package main

import "fmt"

// Assigns value to variable, extension of previous exercise.
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println("s = ", s)
}
