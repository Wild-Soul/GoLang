package main

import "fmt"

func main() {

	defer foo()
	bar()

}

func foo() {
	fmt.Println("Foo has been deferred")
}

func bar() {
	fmt.Println("Bar is 1st")
}
