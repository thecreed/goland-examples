package main

import "fmt"

func varyfunc(numbers ...int64) {
	fmt.Printf("debug: %#v \n", numbers)
}

func verysmartfunc(strs ...string) {
	fmt.Printf("debug: %#v \n", strs)
}

func main() {
	slc := []int64{11, 12, 44, 55, 12}
	varyfunc(slc...)
	varyfunc(222, 322, 22, 22)

	mymojos := []string{"moko", "aaa", "eee", "Ggg"}
	verysmartfunc(mymojos...)

}
