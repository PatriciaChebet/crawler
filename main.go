package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	cliArguments := os.Args[1:]

	if len(cliArguments) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(cliArguments) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %s", os.Args[1])
		htmlBody, err := getHTML(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(htmlBody)
	}
}

func getHTML(rawURL string) (string, error) {
	webpage, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("got Network error: %v", err)
	}
	defer webpage.Body.Close()

	if webpage.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s", webpage.Status)
	}

	if webpage.Header.Get("Content-Type") != "text/html" {
		return "", err
	}

	bodyBytes, err := io.ReadAll(webpage.Body)
	if err != nil {
		return "", err
	}

	htmlBody := string(bodyBytes)

	return htmlBody, nil
}
