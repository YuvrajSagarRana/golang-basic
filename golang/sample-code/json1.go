package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last string
}

func main()  {
	p1 := person{
		First: "Sagar", Last: "Rana",
	}

	p2 := person{
		First:"Hello", Last:"World"}
	xp := [] person {p1, p2}
	fmt.Println(xp)

	bs, error := json.Marshal(xp)
	if error != nil {
		log.Fatalln("Something went wrong while converting")
	}
	fmt.Printf("%+v\n", string(bs))

	// unmarshalling
	data :=`[{"First":"Sagar","Last":"Rana"},{"First":"Hello","Last":"World"}]`
	sliceOfP := [] person {}

	er := json.Unmarshal([] byte(data), &sliceOfP)
	if er != nil {
		log.Fatalln("%v", er)
	}
	fmt.Printf("%v", sliceOfP)


}
