package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
func main()  {

	// get count no of thread
	fmt.Printf("Number of thread %v\n", runtime.GOMAXPROCS(-1))

	for i:=0; i<10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello()  {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment()  {
	counter++
	wg.Done()
}

