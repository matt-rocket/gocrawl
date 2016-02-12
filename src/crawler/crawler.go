package crawler


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"golang.org/x/net/html"
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

	z := html.NewTokenizer(strings.NewReader(page))

	hrefs := make([]string, 10)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return hrefs
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				// we found a link
				attributes := t.Attr
				for _, attr := range attributes {
					if attr.Key == "href" {
						href := attr.Val
						hrefs = append(hrefs, href)
					}
				}
			}
		}
	}
}


func Parser(pageCh chan string, urlCh chan string) {
	for {
		page := <- pageCh
		fmt.Println(extractLinkUrls(page))
	}
}
