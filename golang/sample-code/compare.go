package main

import "fmt"

func main() {
	// go will always convert typeless integer literal or constant or required type
	// implicitly
	faster := 100 > 10.5

	fmt.Println(" faster value " + faster)
}
