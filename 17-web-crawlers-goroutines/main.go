package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(channel chan page, url string) {
	fmt.Println("Getting URL ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- page{url: url, size: len(body)}
}

type page struct {
	size int
	url  string
}

func main() {

	channel := make(chan page)
	//	urls := []string{"http://usa.com", "http://redhat.com", "http://pornhub.com"}
	urlsm := map[string]string{"usa": "http://usa.com"}

	for _, url := range urlsm {
		go responseSize(channel, url)
	}

	for _, url := range urlsm {
		fmt.Println(url, <-channel)
	}

}
