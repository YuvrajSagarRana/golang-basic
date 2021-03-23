package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]
	switch city {

	// No, default behaviour of failing through but can reproduce similar to this
	case "Nepal", "China":
		fmt.Printf("Match the country with %s", city)
	case "India":
		fmt.Printf("Is neighour country")
	default:
		fmt.Printf("None of the city")
	}

}


// two types: express and type
