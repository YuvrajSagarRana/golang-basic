package main

import "fmt"

func main()  {
	person := make(map[string] string)
	person["Name"] = "Sagar"
	person["SurName"] = "Rana"

	if val, exist := person["Name"]; exist {
		fmt.Printf(val)
	}
	
}
