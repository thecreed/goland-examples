package main

import (
	"fmt"
	"testing"
)

func Test_Crawl(t *testing.T) {
	channel := make(chan page)
	urlsm := map[string]string{"usa": "http://usa.com"}

	for _, url := range urlsm {
		go responseSize(channel, url)
	}

	for _, url := range urlsm {
		fmt.Println(url, <-channel)
	}

	//t.Error("this is just a placeholder for a test")
}
