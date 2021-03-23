package main

import "fmt"

func main()  {
	// defer -> executing later
	//fmt.Println("First")
	//defer fmt.Println("Third")
	//fmt.Println("Second")
	//
	// Execute in last in first execute
	defer fmt.Println("First")
	defer fmt.Println("Third")
	defer fmt.Println("Second")


}
