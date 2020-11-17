package dating

import (
	"errors"
	"fmt"
)

// struct type
type Candidate struct {
	Age              int
	Fname            string
	Lname            string
	Gender           string
	Bio              bool
	PreferredPartner string
	ContactInfo
	Address
}

type ContactInfo struct {
	PhoneNumber string
	Email       string
}

type Address struct {
	country     string
	city        string
	street      string
	houseNumber int
}

func (a Address) GetAddress() string {
	return fmt.Sprintf("country: %s , city: %s", a.country, a.city)
}

func (a Address) Country() string {
	return a.country
}

func (a Address) SetCountry(c string) {
	a.country = c
}

func (a Address) City() string {
	return a.city
}

func (a *Address) SetCity(c string) error {
	allowedCities := map[string]int{"Or-Akiva": 1, "Hadera": 1}
	if _, found := allowedCities[c]; found {
		a.city = c
	} else {
		return errors.New(fmt.Sprintf("Couldnt find city %s ", c))

	}

	return nil

}

func (a Address) Street() string {
	return a.street
}

func (a Address) SetStreet(s string) {
	a.street = s
}
