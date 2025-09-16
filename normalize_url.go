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
		return "", err
	}

	return urlObject.Host + urlObject.Path, nil
}
