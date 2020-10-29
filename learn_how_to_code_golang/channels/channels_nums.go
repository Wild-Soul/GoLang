package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	defer func() {
		fmt.Println("ABOUT TO EXIT")
	}()
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for x := range c {
		fmt.Println(x)
	}

}
