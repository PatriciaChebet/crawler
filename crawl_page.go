package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, _ := url.Parse(rawBaseURL)
	hostBaseURL := parsedBaseURL.Host

	parsedCurrentURL, _ := url.Parse(rawCurrentURL)
	hostCurrentURL := parsedCurrentURL.Host

	if hostCurrentURL != hostBaseURL {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatal("Could not normalize URL")
	}

	_, exists := pages[normalizedCurrentURL]
	if exists {
		pages[normalizedCurrentURL]++
	} else {
		pages[normalizedCurrentURL] = 1
	}

	currentURLHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Fatal("Could not get HTML of current url", err)
	}
	fmt.Print(currentURLHTML)

	getsAllURLS, err := GetURLsFromHTML(currentURLHTML, rawBaseURL)
	fmt.Println(getsAllURLS)
	if err != nil {
		log.Fatal("Could not get urls from the HTML body")
	}
	for _, url := range getsAllURLS {
		crawlPage(rawBaseURL, url, pages)
	}

	fmt.Println(pages)

}
