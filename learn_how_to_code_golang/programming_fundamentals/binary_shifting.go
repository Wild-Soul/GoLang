package main

import "fmt"

func main() {

	var a int = 132
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	a = a << 3 // shifts bits by 3 positions to the left.

	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

}
