package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Please provide valide integer number")
	}
	switch  true {
	case i>0:
		fmt.Println("It is positive value")
	case i<0:
		fmt.Printf("It is negative value")
	default:
		fmt.Printf("It is equal to zero")
	}
}

