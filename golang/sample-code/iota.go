package main

import (
	"fmt"
	"strings"
)

func main() {
	// iota is built-on constant generator which generates ever increasing number
	//const (
	//	monday = 0
	//	tuesday = 1
	//	wednesday = 2
	//	thursday = 3
	//	friday = 4
	//	saturday = 5
	//	sunday = 6
	//)
	//fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	//

	const (
		monday = -(iota + 1)
		// black identifier is used to increase the iota value
		_
		_
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	fmt.Println(strings.Repeat("-", 25))

}
