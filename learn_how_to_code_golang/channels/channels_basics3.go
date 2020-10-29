/*
//Fix the following code.
package main

import (
	"fmt"
)

func main() {
	cr := make(<-chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
*/

package main

import (
	"fmt"
)

//In above code the channel is a receive only channel and we're trying to send a value into it.
func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
