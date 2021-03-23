package main

import "fmt"

type Doctor struct {
	number int
	actorName string
	companions []string
}

func  main()  {
	aDoctor := Doctor{
		actorName: "Sagar",
		companions: []string{
			"Sagar Rana",
			"Jio",
		},
	}

// lets drilled down
	// get actorName from Doctor
	fmt.Println(aDoctor.actorName)

	// get the companions name
	fmt.Printf("Companion Name %v", aDoctor.companions)

	// other things
	bDoctor := Doctor{
		3,
		"sagar",
		[] string {
			"Hello World",
			"Golang",
		},
	}
	fmt.Printf("%v", bDoctor)

	fmt.Println("\n>>>>>>>>>>> Anynonymos function")
	anoDoctor := struct {
		name string
	}{ name: "Go Ram"}
	fmt.Println("Anonymous function "+anoDoctor.name)

	// manipulating value and they are not reference type
	sagar := anoDoctor
	sagar.name = "Sagar"
	fmt.Println(sagar.name)
}
