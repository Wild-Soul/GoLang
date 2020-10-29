package main

import "fmt"

func main() {
	cod := struct {
		name   string
		squad  string
		skills []string
	}{
		name:  "Cpt. Price",
		squad: "Bravo Team",
		skills: []string{
			"assassination",
			"sniping",
		},
	}
	fmt.Println(cod)
}
