package main

import "fmt"

func main()  {
	var value interface{} = [...] int {1, 2, 3}
	switch value.(type) {
	case int:
		fmt.Printf("It is integer type. Value is %v", value)
	case float64:
		fmt.Printf("It is float type. Value is %v", value)
	case string:
		fmt.Printf("It is string type. Value is %v", value)
	default:
		fmt.Printf("It is unknown type. Value is %v", value)
	}
}
