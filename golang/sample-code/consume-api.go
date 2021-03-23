package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	response, err := http.Get(`https://api.coinbase.com/v2/prices/spot?currency=USD`)
	if err!=nil {
		log.Fatalln("Fail to fetch the data ", err)
	}

	// read the response using ioutil.ReadAll
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%v", string(data))


	jsonData := map[string] string {"firstname": "Sagar", "LastName": "Rana"}

	marData, err := json.MarshalIndent(jsonData, "", "  ")
	fmt.Printf("\n%v", string(marData))
}
