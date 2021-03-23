package main

import "fmt"

func main() {
	type (
		// Gram underlying type is int
		Gram int

		// Kilogram underlying type is int
		Kilogram Gram

		// Kilogram underlying type is int
		Ton Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)
	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Kilogram(int(apples)))
	fmt.Printf("%T weights package", weights.Gram)

}
