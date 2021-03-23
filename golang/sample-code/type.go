package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	// declare new type
	//type Duration int64;
	//const (
	//	Nanosecond Duration = 1
	//	Microsecond         = Nanosecond * 1000
	//)
	arg := os.Args[1]
	feet, error := strconv.ParseFloat(arg, 64)

	fmt.Printf("%g feet, %g sagar", feet, error)
}
