package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex

	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			var x int = counter
			x = x + 1
			counter = x
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
