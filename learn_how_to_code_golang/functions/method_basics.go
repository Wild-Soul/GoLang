package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi!, I am", p.first, p.last)
	fmt.Println("I'm ", p.age, "years old")
}

func main() {

	p := person{
		first: "Son",
		last:  "Goku",
		age:   32,
	}
	p.speak()

}
