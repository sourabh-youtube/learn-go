package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	c := make(chan int)
	go func() {
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		c <- sum
	}()

	// blocking
	out := <-c

	wg.Wait()
	fmt.Println(out)
}
