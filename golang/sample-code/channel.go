package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main()  {
	// channel is use with go routine
	ch := make(chan int, 40)
	// channel is strongly typed
	wg.Add(2)

	// to make receive only channel
	go func(ch <- chan int) {
		i := <- ch
		fmt.Println(i)

		i = <- ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan <- int) {
		ch <- 42
		fmt.Println("Sender is here")
		ch <- 25
		fmt.Println("Sender is here")
		wg.Done()
	}(ch)
	wg.Wait()
}
