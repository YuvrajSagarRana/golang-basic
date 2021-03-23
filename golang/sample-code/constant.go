package main

import "fmt"

func main() {
	// constants are compile time error
	// while var are run time error
	const meters int = 100
	cm := 100
	m := cm / meters
	fmt.Printf("%dcm is equals to %dm\n", cm, m)

	cm = 200
	m = cm / meters
	fmt.Printf("%dm is %dcm", m, cm)

}
