package main

import (
	"fmt"
	"time"
)

func getTime() time.Time {
	now := time.Now()
	return now
}

func greetUser() {
	ntime := getTime()

	fmt.Printf("Welcome user ! \n")
	fmt.Printf("The time now is %02d:%d:%02d",
		ntime.Hour(),
		ntime.Minute(),
		ntime.Second())
}
