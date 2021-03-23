package main

import "fmt"

type person struct {
	firstname, lastname string
	age                 int
	address             string
}

func main()  {
	sagar := person {firstname:"sagar", lastname:"rana", age: 24}
	fmt.Printf("%+v\n", sagar)

	var anjali person
	anjali.lastname = "Garg"
	anjali.firstname = "Anjali"
	anjali.age = 24
	fmt.Printf("%+v", anjali)


}