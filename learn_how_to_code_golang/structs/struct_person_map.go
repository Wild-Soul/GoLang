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

	pe := person{
		first: "Yamcha",
		last:  "--",
		flavours: []string{
			"blueberry",
			"choco chips",
		},
	}
	persons := map[string]person{
		p.first:  p,
		pe.first: pe,
	}

	for _, v := range persons {
		fmt.Println(v.first, v.last)
		for _, s := range v.flavours {
			fmt.Printf("%s ", s)
		}
		fmt.Println()
	}
}
