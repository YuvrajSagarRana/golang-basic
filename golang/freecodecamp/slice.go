package main

import "fmt"

func main() {
	// array are fixed type
	// len are calculate the length of array or slices
	// if we assign same array to someother variable then, we can have copy but not slices.

	s:= [] int64 {1, 2, 3, 4, 5, 6}
	// they all are pointing to same index and change in one element effect to all
	s[1] = 0
	s[0] = 33
	b := s[1:3]
fmt.Printf("%v\n", s)
	fmt.Printf("New slice is %v\n", b)

	// lets make a slice
	a := make([]int, 3)
	fmt.Println(a)
	fmt.Printf("Length of slice %v\n", len(a))
	// how much spaces left in slice
	fmt.Printf("Length of capacity %d\n", cap(a))
	// slices are always reference type. So, modified make it sometime hard to debug

	// we can append the slice like
	a = append(a, 1)
	// now length get increase by one but you can also
	a = append(a, a ...)
	// we have to destructure if we want to put as we can't append directly array except integer
	fmt.Printf("After destrcturing and appending %v", a)

}
