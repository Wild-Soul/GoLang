package main

import "fmt"

func main() {
	s := "Marshall Mathers"
	if s == "Marshall Mathers" {
		fmt.Println("Eminem")
	} else if s == "Dwayne Michael Carter" {
		fmt.Println("Lil Wayne")
	} else {
		fmt.Println("------")
	}
}
