package main

import (
	"time"
	"crawler"
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

	urlCh := make(chan string)

	pageCh := make(chan string)

	go crawler.SeedUrlMaker(urlCh)

	go crawler.Getter(urlCh, pageCh)

	go crawler.Parser(pageCh, urlCh)

	time.Sleep(time.Millisecond * 10000)
}

