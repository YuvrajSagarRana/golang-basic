package main

import "fmt"

func main() {
	fmt.Println("Sagar is here for freecodecamp.org")
	// array short hand declaration
	grades := [...] int {10, 20}

	cAG := grades
	cAG[0] = 030000
	fmt.Printf("Copy Array %v\n", cAG)
	fmt.Printf("Original Array %v", grades)
	fmt.Printf("\nlength of array %d\n", len(grades))


	// identity matrix
	//var identityMatrix [3][3]int = [3][3]int {[3]int{1, 0, 0 }, [3] int {0,1 ,0}, [3] int {0, 0, 1}}
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0 }
	identityMatrix[1] = [3] int {0,1 ,0}
	identityMatrix[2] =  [3] int {0, 0, 1}
	fmt.Printf("%v", identityMatrix)
}
