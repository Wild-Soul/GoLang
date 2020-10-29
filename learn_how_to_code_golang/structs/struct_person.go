package main

import "fmt"

type person struct {
	first    string
	last     string
	flavours []string
}

func main() {

	p := person{
		first: "Son",
		last:  "Goku",
		flavours: []string{
			"vanilla",
			"chocolate",
			"strawberry",
		},
	}
	fmt.Println(p.first, p.last)
	for i, v := range p.flavours {
		fmt.Println(i, v)
	}

	pe := person{
		first: "Yamcha",
		last:  "--",
		flavours: []string{
			"blueberry",
			"choco chips",
		},
	}
	fmt.Println(pe.first, pe.last)
	for i, v := range pe.flavours {
		fmt.Println(i, v)
	}
}
