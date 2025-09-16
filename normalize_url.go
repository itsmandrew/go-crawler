package main

import (
	"errors"
	"net/url"
)

func normalizeUrl(urlString string) (string, error) {

	if urlString == "" {
		return "", errors.New("url is empty")
	}

	urlObject, err := url.Parse(urlString)

	if err != nil {
		return "", errors.New("invalid url")
	}

	if urlObject.Host == "" {
		return "", errors.New("invalid url")
	}

	return urlObject.Host + urlObject.Path, nil
}
