package main

import (
	"fmt"
)

/*
	//Fix this code
	func main() {
		c:=make(chan int)
		c <- 42
		fmt.Println(<-c)
	}
*/

//Code fix using anonymous functions.
func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

/*
//Code fix using buffered channels.
func main() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}
*/
