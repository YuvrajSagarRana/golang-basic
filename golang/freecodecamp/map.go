package main

import "fmt"

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Calfornia": 39999003,
		"Texas":     2300000,
		"Florida":   23230000,
		"New York":  19999999,
		"Ohio":      2000299,
	}
	statePopulations["Georgia"] = 34343434
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])

	// If key was not found in our map then, ok will print to false
	value, ok := statePopulations["sagar"]
	fmt.Println(value, ok)

	// is it also reference type
	sp := statePopulations
	// to delete the map
	delete(sp, "Ohio") // it will gone delete the key sp from Ohio too
	fmt.Printf("%v", statePopulations)
}
