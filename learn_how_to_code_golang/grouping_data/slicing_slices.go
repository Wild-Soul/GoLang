package main

import "fmt"

func main() {
	ai := []int{}
	for i := 0; i < 100; i++ {
		ai = append(ai, i)
	}

	fmt.Println(ai[42:47])
	fmt.Println(ai[47:52])
	fmt.Println(ai[44:49])
	fmt.Println(ai[43:48])
}
