package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title string `json:"title"`
	Author Author `json:"author"'`
}

type SensorReading struct {
	Name string `json:"name"`
	Capacity int `"json:capacity"`
	Time string `json:"time"`
}

type Author struct {
	Name string `json:"name"`
	Age int 	`json:"age"`
	Developer bool `json:"is_developer"`
}
func main()  {
	author := Author {Name: "Sagar", Age: 23, Developer:true}
	book := Book{ Title: "The lonely Road", Author:author}
	fmt.Printf("%v\n", book)

	jsonString := `{"name":Mobile","capacity":32, "time": "2018-09-23"}`

	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalln("Error converting a struct to json")
	}
	fmt.Printf(string(byteArray))
	}

