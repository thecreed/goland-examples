package main

import (
	"fmt"
	"log"
	"os"
	"utils"
)

//this is incorrect example using two slices instead of a map
//more code + nested for
func countNumOfIdentLinesBad(lines []string) ([]string, []int64) {
	var names []string
	var nameOccures []int64

	for _, val := range lines {

		var exists bool
		//check if the value exists in names
		for idx, valName := range names {
			if val == valName {
				nameOccures[idx]++
				exists = true
			}
		}
		if exists == false {
			names = append(names, val)
			nameOccures = append(nameOccures, 1)
		}

	}
	return names, nameOccures
}

//count number of identical lines in a given slice
//returns a maps
func countNumOfIdentLines(lines []string) map[string]int {

	votes := map[string]int{}
	for _, val := range lines {

		//check if a key exists in a map
		//map can return optional boolean indicating if key is missing
		if count, ok := votes[val]; ok {
			fmt.Printf("exists %s  - %d\n", val, count)
		} else {
			fmt.Printf("Creating new record for %s \n", val)
		}

		votes[val]++
	}
	return votes
}

func main() {

	fname := os.Args[1:]
	fmt.Println("opening files", fname)

	//open multiply files that passed from the CMD
	vals, err := utils.FilesToStrings(fname...)
	fmt.Printf("in total got %d lines from the files \n", len((vals)))
	if err != nil {
		log.Fatal(err)
	}
	//create empty map

	count := countNumOfIdentLines(vals)

	//print the vote count
	for key, val := range count {
		fmt.Printf("line \"%s\" appeared %d times \n", key, val)
	}
	//fmt.Println("count:", count)

}
