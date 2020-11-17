package main

import "fmt"

func changeValue(myInt *int64) {

	*myInt = 12

}

func main() {
	age := int64(37)
	fmt.Printf("value before change %d \n", age)

	changeValue(&age)

	fmt.Printf("value after change %d \n", age)

}
