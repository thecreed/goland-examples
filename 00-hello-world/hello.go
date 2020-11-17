package main

import "fmt"

//DESCR - program tutorial
const DESCR = "Welcome to the golang tutorial"

func printIntro() {
	fmt.Println(DESCR)

}

func main() {
	var id int
	var name string

	printIntro()
	fmt.Printf("your id is: %d \nyour name is %s ", id, name)

}
