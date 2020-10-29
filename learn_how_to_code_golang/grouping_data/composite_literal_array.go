package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", a)
}
