package crawler


import (
	"fmt"
	"os"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
	"golang.org/x/net/html"
)


//func handler(w http.ResponseWriter, r *http.Request) {
//
//}


func SeedUrlMaker(urlCh chan *url.URL) {
	urls := []string{"http://www.google.com", "http://www.dbc.dk", "http://www.nets.dk", "http://www.dr.dk"}

	for _, seedUrl := range urls {
		u, err := url.Parse(seedUrl)
		if err != nil {
			fmt.Println("Could not parse seed url")
			os.Exit(1)
		} else {
			urlCh <- u
		}
	}
}


func getHttpResponse(url *url.URL) (*http.Response, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	return resp, nil
}


func Getter(urlCh chan *url.URL, responseCh chan *http.Response) {
	for {
		url := <-urlCh
		resp, err := getHttpResponse(url)
		if err != nil {
			resp = nil
		}
		responseCh <- resp
	}
}



func extractUrls(resp *http.Response) ([]*url.URL, error) {

	urls := make([]*url.URL, 0)

	response := *resp

	// TODO: the next lines can be simplified
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return urls, err
	}
	z := html.NewTokenizer(strings.NewReader(string(body)))

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return urls, nil
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				// we found a link
				attributes := t.Attr
				for _, attr := range attributes {
					if attr.Key == "href" {
						extractedUrl, err := url.Parse(attr.Val)
						if err != nil {
							return nil, err
						}

						if extractedUrl.Host == "" {
							// fill in the blanks if url is relative
							 extractedUrl = response.Request.Host
						}

						fmt.Println(url)

						urls = append(urls, url)
					}
				}
			}
		}
	}
}


func Parser(responseCh chan *http.Response, urlCh chan *url.URL) {
	for {
		resp := <-responseCh
		urls, err := extractUrls(resp)
		if err != nil {
			fmt.Println("Got error while exxtracting urls")
		}
		fmt.Println(urls)
	}
}
