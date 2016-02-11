package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)


//func handler(w http.ResponseWriter, r *http.Request) {
//
//}


func seedUrlMaker(urlCh chan string) {
	urls := []string{"http://www.google.com", "http://www.dbc.dk", "http://www.nets.dk", "http://www.dr.dk"}

	for _, url := range urls {
		urlCh <- url
	}
}


func getPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	page := string(body)

	return page, nil
}


func getter(urlCh chan string, pageCh chan string) {
	for {
		url := <-urlCh
		page, err := getPage(url)
		if err != nil {
			page = ""
		}
		pageCh <- page
	}
}


func parser(pageCh chan string, urlCh chan string) {
	for {
		page := <- pageCh
		fmt.Println(len(page))
	}
}


func main() {
	//	http.HandleFunc("/", handler)
	//	http.ListenAndServe(":8080", nil)

	urlCh := make(chan string)

	pageCh := make(chan string)

	go seedUrlMaker(urlCh)

	go getter(urlCh, pageCh)

	go parser(pageCh, urlCh)

	time.Sleep(time.Millisecond * 3000)
}

