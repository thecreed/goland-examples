// game that let player guess a random generated number from 1 to 100
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func askFor(question string) (string, error) {
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')

	answer = strings.TrimSpace(answer)
	return answer, err
}

func askForInt(question string) (int64, error) {
	guessNum, _ := askFor(question)
	guessNumInt, err := strconv.ParseInt(guessNum, 10, 32)

	return guessNumInt, err
}

func genRandomNumber() int64 {
	epoch := time.Now().Unix()
	rand.Seed(epoch)
	num := rand.Int63n(100) + 1
	return num
}

func main() {
	totalAttempts := 3
	genNumber := genRandomNumber()
	log.Printf("generated number %d", genNumber)

	x := 0
	success := false
	for x = 0; x < totalAttempts; x++ {
		guessNum, _ := askForInt("Enter your Guesed number: ")
		log.Printf("Guess attempt number %d/%d , value: %d", x, totalAttempts, guessNum)
		if guessNum < genNumber {
			log.Printf("you guess is too low")
		} else if guessNum > genNumber {
			log.Println("Your guess is too high")
		} else {
			success = true
			log.Printf("You've got it right after %d attempts!", x)
			break
		}
	}
	if !success {
		log.Println("You didnt guessed it at all :( ")
	}

}
