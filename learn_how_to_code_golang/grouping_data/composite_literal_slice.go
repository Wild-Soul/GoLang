package main

import "fmt"

func main() {
	ai := []int{}
	for i := 0; i < 10; i++ {
		ai = append(ai, i)
	}
	for i, v := range ai {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", ai)
}
