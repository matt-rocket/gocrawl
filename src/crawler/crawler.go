package crawler


import (
	"fmt"
	"net/http"
	"io/ioutil"
)


//func handler(w http.ResponseWriter, r *http.Request) {
//
//}


func SeedUrlMaker(urlCh chan string) {
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


func Getter(urlCh chan string, pageCh chan string) {
	for {
		url := <-urlCh
		page, err := getPage(url)
		if err != nil {
			page = ""
		}
		pageCh <- page
	}
}


func extractLinkUrls(page string) []string {
	return []string{"bla", "blabla"}
}

func Parser(pageCh chan string, urlCh chan string) {
	for {
		page := <- pageCh
		fmt.Println(extractLinkUrls(page))
	}
}
