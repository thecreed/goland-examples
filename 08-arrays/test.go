package main

import (
	"fmt"
	"utils"
)

const numberOfRecords int64 = 4

func printTheRecords(records [numberOfRecords]string) {
	fmt.Printf("%6s | %10s \n", "index", "name")
	fmt.Println("------------------")

	for x, value := range records {
		fmt.Printf("%6d | %10s \n", x, value)
		records[x] = "deleted"
	}

}

func printTheRecordsNice(records [numberOfRecords]string) {

	for x := range records {
		fmt.Printf("%6d \n", x)

	}
}

func printTheRecordsPtr(records *[numberOfRecords]string) {
	fmt.Printf("%6s | %10s \n", "index", "name")
	fmt.Println("------------------")
	for x, value := range records {
		fmt.Printf("%6d | %10s \n", x, value)
		records[x] = "deleted"
	}

}

func main() {
	//numOfElements, _ := utils.AskForInt("Number of records to add?")

	//create array with numOfElements
	var records [numberOfRecords]string

	//collect records

	for x := range records {
		fmt.Printf("hi there - %d \n", x)
		records[x] = utils.AskForString(fmt.Sprintf("enter your record n %d", x))

	}

	printTheRecordsPtr(&records)
	printTheRecords(records)
	printTheRecordsNice(records)

}
