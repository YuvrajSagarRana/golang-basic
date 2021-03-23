package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// nill means there was no error
	// there is no try and catch for exception handling
	// strconv.Itoa this function never return error as it always got successful
	// strconv.Atoi (s string) (int, error)
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Hello there is an error")
	}
	fmt.Println(n, err)
}
