package main

import (
	"log"
	"net/http"
	"time"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("welcome to my awseome site !")
	time.Sleep(time.Second * 3)

	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/welcome", viewHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	//err1 := http.ListenAndServeTLS("localhost:8181", nil)
	log.Fatal(err)

}
