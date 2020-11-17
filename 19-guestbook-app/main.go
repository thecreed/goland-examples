package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const datafile string = "data/signatures.txt"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type signatures struct {
	Values []string
	Count  int
}

//read text file with data, return slice with lines
func getStrings(fileName string) *signatures {
	//var lines []string
	sigs := signatures{}
	file, err := os.Open(fileName)

	if os.IsNotExist(err) {
		return nil
	}
	checkError(err)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sigs.Values = append(sigs.Values, scanner.Text())
	}
	sigs.Count = len(sigs.Values)

	checkError(scanner.Err())

	return &sigs
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {

	lst := getStrings(datafile)

	log.Println(lst)
	html, err := template.ParseFiles("index.html")
	checkError(err)
	//	err = html.Execute(os.Stdout, lst)
	//checkError(err)
	err = html.Execute(writer, lst)
	checkError(err)
}

func appendNewLineToFile(fileName string, newLine string) {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(fileName, options, os.FileMode(0600))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(file, newLine)

}

func viewHandlerNewSig(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" && request.FormValue("sig_body") != "" {
		fmt.Println("sig_body", request.FormValue("sig_body"))
		appendNewLineToFile(datafile, request.FormValue("sig_body"))
		http.Redirect(writer, request, "/guestbook", http.StatusFound)

	}

	html, err := template.ParseFiles("new.html")
	checkError(err)
	err = html.Execute(writer, nil)
	checkError(err)
}

func redirectGuestBook(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/guestbook", 301)
}

func main() {
	//http.HandleFunc("/", redirectGuestBook)
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", viewHandlerNewSig)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
