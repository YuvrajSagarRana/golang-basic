package main

import "fmt"

func main()  {
	// array definition
	// first have size [3] int general: [size] datatype { "", ""}
	grades := [3] int {30, 3, 23}
	fmt.Printf("%v", grades)

	// incase of literal syntax we can have
	gradesS := [...] int {20, 30, 309}
	fmt.Println(gradesS)

	// define array
	var students[3] string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Sagar Rana"
	students[1] = "is"
	students[2] = "Testing"
	fmt.Println(students)

	// length of string
	fmt.Println(len(students[0]))
	fmt.Printf("Length of array students %T\n", students)

	// Identity Matrix
	var identityMatrix [3][3]int = [3][3]int {[3]int {1, 0, 0}, [3] int {0, 1, 0}, [3]int {0, 0, 1}}
	fmt.Printf("%v\n", identityMatrix)

	// identity matrix of 2*2
	var identityMatrix2 [2][2] int = [2][2]int {[2] int {1,0}, [2] int {0, 1}}
	fmt.Printf("%v", identityMatrix2)

	// we can declare arrays of arrays by assigning value between them like this:
	var identityMatrix3 [3][3] int;
	identityMatrix3[0] = [3] int {1, 0, 0}
	identityMatrix3[1] = [3] int {0, 1, 0}
	identityMatrix3[2] = [3] int {0, 0, 1}
	fmt.Printf("\n%v", identityMatrix3)

	// Array means actually copying value not reference when assigning to new variable
	// slice syntax
	a := []int {1, 2, 3} // slice example
	fmt.Printf("\n%d", len(a))
	a[0] = 3
	a[2]= 3
	fmt.Printf("\nCapacity %d", cap(a))

	// for creating a slice we are using built in function
	slice1 := make([]int, 3, 100) // us
	// ing this function we can get intialize the value like [0, 0, 0]
	// 100 is capcity


}
