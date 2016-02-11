package main

import (
	"time"
	"crawler"
)

func main() {
	//	http.HandleFunc("/", handler)
	//	http.ListenAndServe(":8080", nil)

	urlCh := make(chan string)

	pageCh := make(chan string)

	go crawler.SeedUrlMaker(urlCh)

	go crawler.Getter(urlCh, pageCh)

	go crawler.Parser(pageCh, urlCh)

	time.Sleep(time.Millisecond * 10000)
}

