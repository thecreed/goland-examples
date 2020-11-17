package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pVersion = 0.3

const greetingmsg string = `
Welcome to the fileInfo program,
it will prompt you to enter a filename ,
and it will print  various information about that file.

`

func getFileInfo(path string) int64 {
	fmt.Printf("Getting info about path: %s", path)

	finfo, err := os.Stat(path)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file \nsize: %d\nperm: %s\n", finfo.Size(), finfo.Mode())
	return finfo.Size()
}

func askFor(question string) (string, error) {
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')

	answer = strings.TrimSpace(answer)
	return answer, err
}

func askForMinSize() int64 {
	minSize, err := askFor("Press enter size of file (integer):")

	if err != nil {
		fmt.Println(err)
	}

	minSizeInt, _ := strconv.ParseInt(minSize, 10, 32)

	return minSizeInt

}
func askForFileName() string {
	path, err := askFor("Press enter a with with its full path")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Entered path %s\n", path)

	return path
}
func main() {
	fmt.Printf(" %s", greetingmsg)
	path := askForFileName()
	minSize := askForMinSize()
	fileSize := getFileInfo(path)

	status := 0
	if fileSize < minSize {
		fmt.Printf("file %s size is less then %d", path, minSize)
		status = 2
	}
	fmt.Printf("the status is %d", status)
}
