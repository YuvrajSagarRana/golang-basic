package main

import (
	"fmt"
	"sync"
)

// to synchronize multiple go routine
var wg = sync.WaitGroup{}

func main(){
	msg := "Hello"
	wg.Add(1)
   go func(msg string) {
   	fmt.Printf("%s", msg)
   	wg.Done()
   }(msg)
	msg = "GoodBye"
   wg.Wait()
}

func sayHello(){
	fmt.Printf("Hello go routine\n")
}
