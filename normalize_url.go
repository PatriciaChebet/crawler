package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	host := parsedURL.Host
	path := parsedURL.Path

	trimmedPath, _ := strings.CutSuffix(path, "/")

	expectedParsedURL := host + trimmedPath

	return expectedParsedURL, nil
}
