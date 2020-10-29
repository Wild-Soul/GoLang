package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//There are other ways to do the same thing. But I'm gonna stick with anonymous function.
	wg.Add(1)
	go func() {
		defer wg.Done()
		foo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bar()
	}()

	wg.Wait()
}

func foo() {
	fmt.Println("FOO")
}

func bar() {
	fmt.Println("BAR")
}
