package main

import "fmt"

func main() {

	fmt.Println(foo())
	fmt.Println(bar())

}

func foo() int {
	return 1024
}

func bar() (int, string) {
	name := "Kakarot"
	return len(name), name
}
