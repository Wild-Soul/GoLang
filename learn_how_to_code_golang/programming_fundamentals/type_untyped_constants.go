package main

import "fmt"

const (
	a     = 42
	b int = 54
)

func main() {

	fmt.Printf("%v\t%v\n", a, b)
	fmt.Printf("%T\t%T", a, b)

}
