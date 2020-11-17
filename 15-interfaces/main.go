package main

import (
	"fmt"
	"log"
	"time"
)

type person struct {
	name       string
	age        int
	gender     string
	birthday   time.Time
	is_married bool
}

type employee struct {
	person
	salary      int
	startedWork time.Time
}

func (p person) greetBirthday() {
	fmt.Printf("Hi %s, happy birthday to to , you are %d years old today \n", p.name, p.age+1)
	fmt.Printf("your birthday at %s\n", p.birthday.Format("Jan 2, 2006"))

}

func (e employee) greetBirthday() {
	fmt.Printf("Hi Employee %s , you are entitles to birthday gift bonus , you'r next salary will be %d ! \n", e.name, e.salary+1000)
	fmt.Println("you started to work at", e.startedWork)
	fmt.Printf("your birthday at %s\n", e.birthday.Format("Jan 2, 2006"))
}

type member interface {
	greetBirthday()
}

func greetUser(s member) {
	s.greetBirthday()

	fmt.Println("DEBUG", s)
	//try to do interface type assertion
	buser, ok := s.(person)
	if ok && buser.is_married {
		fmt.Println("its time to consider devorce !")
	}
}

func main() {
	fmt.Println("Welcome to my intefaces tester program !")
	emp := employee{salary: 2100, startedWork: time.Now()}
	emp.name = "Jhonny"
	emp.age = 36
	// This is the value we are trying to parse.
	value := "1983-05-28"

	// The form must be January 2,2006.
	form := "2006-01-02"

	mydate, err := time.Parse(form, value)
	if err != nil {

		log.Fatal(err)
	}
	emp.birthday = mydate
	per := person{age: 25, name: "michael", is_married: true}
	per.birthday = time.Date(1983, 5, 28, 12, 0, 0, 0, time.UTC)
	//fmt.Println("date:", kkk)
	//emp.greetBirthday()
	//per.greetBirthday()
	//	var mem member = emp
	greetUser(emp)
	greetUser(per)
	//greetUser(mem)

}
