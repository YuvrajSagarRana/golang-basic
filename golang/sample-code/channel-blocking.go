package main

import (
	"fmt"
	"sync"

)

var wg sync.WaitGroup
func main()  {
	c := boring("boring")

	for i:=0;i<5;i++ {
		fmt.Printf("Before first run here\n")
		fmt.Printf("You say %q\n", <-c)
		fmt.Printf("Last run here\n")
	}
	wg.Wait()
	fmt.Printf("Program exited")
}

func boring(msg string) <- chan string {
	c := make(chan string)
	wg.Add(1)
	go func() {
		for i:=0; ;i++ {
			fmt.Printf("First run here\n")
			c <- fmt.Sprintf("%s %d", msg, i)
			fmt.Printf("Second run here\n")
			//time.Sleep(2 * time.Millisecond)
		}
	}()
	wg.Done()

	return c
}
