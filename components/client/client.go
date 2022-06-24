package client

import "net/http"

// Search username from source
func GetUrl(url string) (*http.Response, error) {
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
