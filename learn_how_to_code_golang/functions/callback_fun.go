package main

import "fmt"

func sum(ai []int) int {
	ans := 0
	for _, i := range ai {
		ans = ans + i
	}
	return ans
}

func main() {

	ai := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(even(sum, ai))
}

func even(sum func(ai []int) int, ai []int) int {
	var xi []int
	for _, i := range ai {
		if i%2 == 0 {
			xi = append(xi, i)
		}
	}
	return sum(xi)
}
