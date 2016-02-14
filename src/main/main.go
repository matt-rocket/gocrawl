package main

import (
	"time"
	"crawler"
	"net/http"
	"net/url"
)

type persons []int

func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	//	http.HandleFunc("/", handler)
	//	http.ListenAndServe(":8080", nil)

	urlCh := make(chan *url.URL)

	responseCh := make(chan *http.Response)

	go crawler.SeedUrlLoader(urlCh)

	go crawler.Getter(urlCh, responseCh)

	go crawler.Parser(responseCh, urlCh)

	time.Sleep(time.Millisecond * 10000)
}
