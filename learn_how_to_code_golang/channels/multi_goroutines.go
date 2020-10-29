package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			go func() {
				for j := 0; j < 10; j++ {
					c <- j
				}
			}()
		}
		//close(c)
	}()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(<-c)
		}
	}

	fmt.Println("About to exit")
}
