package main

import "fmt"

func main() {

	a := 10
	fmt.Println(&a)
	foo(&a)
	fmt.Println(a)

}

func foo(a *int) {
	*a = 100
	fmt.Printf("%T\t,%v\n", a, a)
}
