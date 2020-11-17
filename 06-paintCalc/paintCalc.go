package main

import (
	"fmt"
	"log"
)

func calcNeededPaint(width float64, height float64, paintCoverage float64) (float64, error) {
	var err error
	if width <= 0 {
		err = fmt.Errorf("width: value  %f is invalid", width)
	}

	if height <= 0 {
		err = fmt.Errorf("height: value %f is invalid", height)
	}

	if err != nil {
		return 0, err
	}

	roomDimensions := width * height
	neededPaint := roomDimensions / paintCoverage
	return neededPaint, err
}

func main() {

	literOfPaint, err := calcNeededPaint(20, -10.3, 10)

	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Printf("for living room we need %.4f liters of paint", literOfPaint)

}
