package main

import (
	"fmt"
	"math"
	"strings"
)

func maximum(vals ...float64) float64 {
	max := math.Inf(-1)

	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max

}

func makeString(parm ...string) {

	fmt.Printf("dumping: %#v \n", parm)

	bigstr := strings.Join(parm, "#-#")

	fmt.Println("bigstr=", bigstr)
	fmt.Println(true, "sdsada dsa", 23232)

}

func main() {

	slice := []string{"bobo",
		"rambo",
		"stalona",
		"vandaom"}

	slicepre := make([]string, 100)

	var emptyslice []string

	fmt.Printf("%#v \n\n", slice)
	fmt.Printf("%#v \n\n", emptyslice)
	fmt.Printf("%#v \n\n", slicepre)

	makeString("bobo", "kakaka", "royroy")

	val := maximum(22.33, 23.224, 23.42)
	fmt.Println(val)
}
