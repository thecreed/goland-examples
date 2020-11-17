package main

import "fmt"

type Age int

func (a Age) Say() {
	fmt.Println("My age is ", a)

}

func main() {
	var joeAge Age
	joeAge = 12

	dinnaAge := Age(18)
	dinnaAge.Say()
	joeAge.Say()

}
