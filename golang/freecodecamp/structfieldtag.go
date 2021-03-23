package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string `required max:"100"`
	Origin string `Nothing required`
}
func main() {
	// to get that tag we need to use reflection package
	t := reflect.TypeOf(Animal{})
	field, ret := t.FieldByName("Name")
	fmt.Println(field, ret)
}
