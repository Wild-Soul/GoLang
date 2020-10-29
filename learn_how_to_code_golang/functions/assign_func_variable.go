package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Assigned a function to a variable and then called it")
	}

	f()
	fmt.Printf("%T\n", f)
}
