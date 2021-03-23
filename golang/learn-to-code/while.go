package main

import (
	"fmt"
	"strconv"
)

func main()  {
	i := 20
	for i> 10 {
		fmt.Printf("i value "+strconv.Itoa(i))
		i++
		if(i == 30) {
			break
		}
	}

	for {
		// infinite loop
		i++
	}
}
