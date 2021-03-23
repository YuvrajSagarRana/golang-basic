package main

import (
	"fmt"
	"sync"
)

// create a waitGroup from sync package to wait for the other routine
var ( wg sync.WaitGroup )
func main()  {
	// create a channel
	done := make(chan bool)
	wg.Add(1)
	go func() {
		done <- true
		fmt.Printf("first it run\n")
		defer wg.Done()
	}()
	fmt.Printf("It come here and get blocked\n")
	fmt.Println(<-done)
	fmt.Printf("Get unblocked")
	wg.Wait()
}
