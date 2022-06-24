package service

import (
	"fmt"
	"net/http"

	"github.com/izzxx/username_checker/components/client"
	"github.com/izzxx/username_checker/components/schema"
	"github.com/izzxx/username_checker/components/utility"
)

// Check "./schema/source.go" to see avaible sources

// Get information from source
func UsernameChecker(req *http.Request, username string) (string, error) {
	requestURI := req.URL.RequestURI()

	response, err := CheckUrlUsername(requestURI)
	if err != nil {
		return "", err
	}

	return response, nil
}

// search users by source
func CheckUrlUsername(url string) (string, error) {
	detail, username, ok := utility.ValidateUrl(url)
	if !ok {
		return "", schema.ErrUrlFormat
	}

	switch detail {
	case schema.Facebook:
		urlReq := fmt.Sprintf("%s%s", schema.FacebookUrl, username)

		url, err := urlChecker(urlReq)
		if err != nil {
			return "", err
		}

		return url, nil
	case schema.Instagram:
		urlReq := fmt.Sprintf("%s%s", schema.InstagramUrl, username)

		url, err := urlChecker(urlReq)
		if err != nil {
			return "", err
		}

		return url, nil
	case schema.Github:
		urlReq := fmt.Sprintf("%s%s", schema.GithubUrl, username)

		url, err := urlChecker(urlReq)
		if err != nil {
			return "", err
		}

		return url, nil
	default:
		return "", schema.ErrSourceNotFound
	}
}

// Check url with username
func urlChecker(url string) (string, error) {
	resp, err := client.GetUrl(url)
	if err != nil {
		return "", schema.ErrSource
	}

	// Response code my have been modified from source,
	// I,m just following the standard
	if resp.StatusCode > 299 {
		return "", schema.ErrUsernameNoMatch
	}

	if resp.StatusCode == 200 {
		return url, nil
	}

	return "", schema.ErrSourceNotFound
}
