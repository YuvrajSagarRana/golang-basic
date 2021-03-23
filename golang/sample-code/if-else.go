package main

import (
	"fmt"
	"strings"
)

func main() {
	speed := false
	if speed  {
		fmt.Println("if-condition")
	} else {
		fmt.Println("else-condition")
	}

	trimValue := strings.TrimSpace("sagar ")
	fmt.Println(len(trimValue))
}
