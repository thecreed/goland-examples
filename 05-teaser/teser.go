package main

import (
	"fmt"
)

func main() {
	mystr := fmt.Sprintf("this is a var test %#v \n", "\t \b \b \t \n ")
	mystr += fmt.Sprintf("%10s | %4s\n", "lalala", "22")
	mystr += fmt.Sprintf("%10s | %4s\n", "lalssssala", "32")
	mystr += fmt.Sprintf("%10d | %4s\n", 12, "32")

	fmt.Println(mystr)
}
