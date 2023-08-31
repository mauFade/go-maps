package main

import "log"

type User struct {
	FirstName  string
	SecondName string
}

func main() {
	myMap := make(map[string]User)

	myUser := &User{
		FirstName:  "Mau",
		SecondName: "Cardoso",
	}

	myMap["me"] = *myUser

	log.Println(myMap["me"].SecondName, myMap["me"].FirstName)
}
