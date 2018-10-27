package main

import(
	"strings"
	"golang.org/x/net/html"
)

func extractDownloadLinks(webPage string) []string {
	var foundUrls = make([]string, 0)
	pageReader := strings.NewReader(webPage)
	
	tokenizedPage := html.NewTokenizer(pageReader)
	
	for {
		analizedToken := tokenizedPage.Next()
		switch {
		case analizedToken == html.ErrorToken:
			// End of the document, we're done
			return foundUrls

		case analizedToken == html.StartTagToken:
			thisToken := tokenizedPage.Token()
			// if is not a link token then skip it
			if thisToken.Data != "a" {
				continue
			}

			// Extract the href value, if there is one
			url, ok := getHref(thisToken)
			if !ok {
				continue
			}

			// Make sure the url begines in http**
			hasProto := strings.Index(url, "/servlet/servlet.OrgExport") >= 0
			if hasProto {
				foundUrls = append(foundUrls, url)		
			}
		}
		
	} 
}

func getHref(token html.Token) (href string, successful bool) {
	
	for _, thisAttribute := range token.Attr {
		if thisAttribute.Key == "href" {
			href = thisAttribute.Val
			successful = true
			break
		}
	}
	
	return
}
