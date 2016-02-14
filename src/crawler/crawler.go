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


func SeedUrlLoader(urlCh chan *url.URL) {
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


func extractPageContent(resp *http.Response) ([]*url.URL, []string, error) {

	urls := make([]*url.URL, 0)

	textLines := make([]string, 0)

	response := *resp

	// TODO: the next lines can be simplified
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return urls, textLines, err
	}
	z := html.NewTokenizer(strings.NewReader(string(body)))

	// keep track of when we enter a script element
	inScriptTag := false

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return urls, textLines, nil
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			isScript := t.Data == "script"
			if isAnchor {
				// we found a link
				attributes := t.Attr
				for _, attr := range attributes {
					if attr.Key == "href" {
						extractedUrl, err := url.Parse(attr.Val)
						if err != nil {
							return nil, textLines, err
						}
						if extractedUrl.Host == "" {
							// fill in the blanks if url is relative
							baseUrl := response.Request.URL
							extractedUrl.Host = baseUrl.Host
							extractedUrl.Scheme = baseUrl.Scheme
						}
						urls = append(urls, extractedUrl)
					}
				}
			} else if isScript {
				inScriptTag = true
			}
		case tt == html.EndTagToken:
			t := z.Token()
			isScript := t.Data == "script"
			if isScript {
				inScriptTag = false
			}
		case tt == html.TextToken:
			t := z.Token()
			line := t.String()
			line = strings.TrimSpace(line)
			if line != "" && !inScriptTag {
				textLines = append(textLines, line)
			}
		}
	}
}


func Parser(responseCh chan *http.Response, urlCh chan *url.URL) {
	for {
		resp := <-responseCh

		urls, _, err := extractPageContent(resp)
		if err != nil {
			fmt.Println("Got error while extracting urls")
		}

		fmt.Println(urls)

		//for i, line := range textLines {
			//fmt.Println(i, line)
		//}
	}
}
