package main

import (
	"dating"
	"fmt"
	"log"
)

func showInfo(user dating.Candidate) {
	fmt.Printf(" user: %5s \n last: %5s \n age: %5d \n gender: %5s \n address: %s",
		user.Fname, user.Lname, user.Age, user.Gender, user.GetAddress())

}

func defaultMember(fname string) dating.Candidate {
	member := dating.Candidate{
		Fname:            fname,
		Gender:           "women",
		PreferredPartner: "women",
		Address:          dating.Address{},
	}

	return member
}
func initStruct() {
	// struct variable
	var myData struct {
		age              int
		fname            string
		lname            string
		gender           string
		bio              bool
		preferredPartner string
	}

	members := []dating.Candidate{}

	joe := defaultMember("joe")
	err := joe.SetCity("Or-Akiva")
	if err != nil {
		log.Fatal(err)
	}
	members = append(members, joe)

	if false {
		fmt.Printf("debug: %#v \n	", myData)
	}

	if true {
		fmt.Printf("debug: %#v \n	", members)
	}
	showInfo(joe)

}

func main() {
	fmt.Println("Welcome to my awesome struct program !")
	initStruct()

}
