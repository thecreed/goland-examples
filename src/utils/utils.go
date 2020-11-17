// Package utils - provides various utilities for the average joe developer
package utils

import (
	"bufio"
	"fmt"
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

//AskForInt - prompt for integet
func AskForInt(question string) (int64, error) {
	guessNum, _ := askFor(question)
	guessNumInt, err := strconv.ParseInt(guessNum, 10, 32)

	return guessNumInt, err
}

//AskForString - prompt for string
func AskForString(question string) string {
	guessNum, _ := askFor(question)

	return guessNum
}

//GenRandomNumber - generate random number between 1 and 100
func GenRandomNumber() int64 {
	epoch := time.Now().Unix()
	rand.Seed(epoch)
	num := rand.Int63n(100) + 1
	return num
}

//FilesToStrings - open one or more files and convert their values in strings slice
func FilesToStrings(fnames ...string) ([]string, error) {
	fmt.Printf("%#v", fnames)

	var lines []string
	for _, fname := range fnames {
		file, err := os.Open(fname)

		if err != nil {
			//log.Fatal(err)
			continue
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
		}
		file.Close()
	}

	//fmt.Printf("%#v", lines)

	return lines, nil

}
