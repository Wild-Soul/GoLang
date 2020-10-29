package main

import "fmt"

func main() {
	born := 1998
	current := 2020
	i := born
	for i <= current {
		fmt.Println(i)
		i++
	}
}
