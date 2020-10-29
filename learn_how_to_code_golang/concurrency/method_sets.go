package main

import (
	"fmt"
)

type person struct {
	name string
	age  int8
}

func (p *person) speak() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "JACK",
		age:  32,
	}
	// Will not work because speak method has pointer reciver
	//saySomething(p)

	//works since i'm passing a pointer and speak method has a pointer receiver
	saySomething(&p)
}
