package main

import "fmt"

type chocos int

func main() {
	var x chocos
	fmt.Printf("%v \n %T\n", x, x)
	x = 42
	fmt.Println(x)
}
