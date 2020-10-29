package main

import "fmt"

type person struct {
	age   int
	first string
	last  string
}

func foo(p *person) {

	(*p).age = 43
	(*p).first = "Vegeta"
	(*p).last = "--"
}

func main() {

	p := person{
		age:   32,
		first: "Son",
		last:  "Goku",
	}
	fmt.Println("Before:", p)
	foo(&p)
	fmt.Println("After:", p)

}
