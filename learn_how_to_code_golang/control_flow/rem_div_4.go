package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d ", i%4)
		if i%5 == 0 {
			fmt.Println()
		}
	}
}
