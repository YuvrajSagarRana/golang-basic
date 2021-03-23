package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func printIfPerson(object interface{}) {
	person, ok := object.(Person)

	if ok {
		fmt.Printf("Hello %s!\n", person.firstName)
	}
}

func main()  {
	printIfPerson()
}
