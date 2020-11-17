package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type groceryCountError int

const dirP string = "/tmp/groceries/"

func (g groceryCountError) Error() string {
	return fmt.Sprintf("error: int %d is invalid grocery count, value must be > 0", g)
}

type groceryList map[string]int64

func printError(err error) {
	if err != nil {
		fmt.Printf("your having error: %s \n", err)
	}
}

// recursive function to parse all the files at subdirs
func (g groceryList) initList(path string) error {
	defer reportPanic()
	//go over the dirP content and check each file
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, finfo := range files {
		filePath := filepath.Join(path, finfo.Name())
		if finfo.IsDir() {
			fmt.Printf("%s is dir,traversing dir \n", finfo.Name())
			err := g.initList(filePath)
			if err != nil {
				log.Println(err)
				//return err
			}

		} else {
			g.openGroceryFile(fmt.Sprint(filePath))

		}

		//	return nil

	}

	//if its a regular text file , then try parse it to the dict
	//g.openGroceryFile()
	return nil
}

func (g groceryList) openGroceryFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")
		valNumber, _ := strconv.ParseInt(lines[1], 10, 32)
		g[lines[0]] = valNumber

		//	content = append(content, scanner.Text())
	}
	return nil
}

func (g groceryList) addItem(item string, count int) error {

	if count <= 0 {
		return groceryCountError(count)
	}

	//g[item] = count

	return nil

}

func reportPanic() {
	p := recover()

	if p == nil {
		return
	}

	err, ok := p.(error)

	if ok {
		fmt.Println(err)
	}

}

func main() {
	mylist := groceryList{"milk": 2}
	defer reportPanic()
	err := mylist.initList(dirP)

	if err != nil {
		log.Fatal(err)
	}

	err = mylist.addItem("eggs", 2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mylist)
}
