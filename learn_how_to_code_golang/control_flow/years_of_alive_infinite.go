package main

import "fmt"

func main() {
	born := 1998
	current := 2020
	i := born
	for {
		if i > current {
			break
		}
		fmt.Println(i)
		i++
	}
}
