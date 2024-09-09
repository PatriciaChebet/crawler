package main

import (
	"fmt"
	"io"
	"net/http"
)

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
