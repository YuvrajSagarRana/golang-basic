package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var ( n int
	      err error)
	if a:= os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		fmt.Printf("The value of n is %v\n", n)
	}
	fmt.Printf("n is %d - you have been shadowed it", n)
}
