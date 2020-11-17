package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const fileName string = "/etc/hosts"

func readFileContent(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	var content []string
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return content, nil
}

func main() {

	args := os.Args[1:]
	fmt.Println("Passed args: ", args)

	content, _ := readFileContent(args[0])
	fmt.Printf("dumping  content of \"%s\" \n", fileName)

	for lineNum, line := range content {
		fmt.Printf("line %d: %s\n", lineNum+1, line)
	}

	contentSlices := content[1:3]
	contentSlices[0] = "deleted"
	for lineNum, line := range content {
		fmt.Printf("line %d: %s\n", lineNum+1, line)
	}

}
