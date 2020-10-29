package main

import "fmt"

func main() {

	ai := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(foo(ai...))
	fmt.Println(bar(ai))
}

func foo(ai ...int) int {
	ans := 0
	for i := 0; i < len(ai); i++ {
		ans = ans + ai[i]
	}
	return ans
}

func bar(ai []int) int {
	sum := 0
	for i := 0; i < len(ai); i++ {
		sum = sum + ai[i]
	}
	return sum
}
