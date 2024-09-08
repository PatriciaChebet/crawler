package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromHTML(body string, baseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return []string{}, err
	}
	relativeURLs := treeTraversal(doc)
	absoluteURLs, err := convertsToAbsoluteURLs(relativeURLs, baseURL)
	if err != nil {
		return []string{}, err
	}
	return absoluteURLs, nil
}

func treeTraversal(treeNode *html.Node) []string {
	urls := []string{}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					urls = append(urls, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(treeNode)
	return urls
}

func convertsToAbsoluteURLs(relativeURLs []string, baseURL string) ([]string, error) {
	absoluteURLs := []string{}
	for _, relativeURL := range relativeURLs {
		parsedURL, err := url.Parse(relativeURL)
		if err != nil {
			return []string{}, err
		}
		if parsedURL.IsAbs() {
			absoluteURLs = append(absoluteURLs, relativeURL)
		} else {
			absoluteURL := baseURL + relativeURL
			absoluteURLs = append(absoluteURLs, absoluteURL)
		}
	}
	return absoluteURLs, nil
}
