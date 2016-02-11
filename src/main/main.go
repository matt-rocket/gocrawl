package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


//func handler(w http.ResponseWriter, r *http.Request) {
//
//}


func urlMaker(urlCh chan string) {
	urls := []string{"http://www.google.com", "http://www.dbc.dk", "http://www.nets.dk", "http://www.dr.dk"}


	i := 0
	for {
		url := urls[i]
		i += 1
		if i >= len(urls){
			i = 0
		}
		urlCh <- url
	}

	close(urlCh)
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
		url := <- urlCh
		page, err := getPage(url)
		if err != nil {
			page = ""
		}
		pageCh <- page
	}
}


func main() {
	//	http.HandleFunc("/", handler)
	//	http.ListenAndServe(":8080", nil)

	urlCh := make(chan string)

	pageCh := make(chan string)

	go urlMaker(urlCh)

	go getter(urlCh, pageCh)

	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))
	fmt.Println(len(<-pageCh))

}