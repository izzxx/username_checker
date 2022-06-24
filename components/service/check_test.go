package service

import (
	"testing"

	"github.com/izzxx/username_checker/components/schema"
)

func TestUrlChecker(t *testing.T) {
	var urls = []struct {
		order       string
		description string
		url         string
	}{
		{
			order:       "1",
			description: "success",
			url:         "https://github.com/izzxx",
		},
		{
			order:       "2",
			description: "success",
			url:         "https://instagram.com/jokowi",
		},
		{
			order:       "3",
			description: "success",
			url:         "https://facebook.com/jokowi",
		},
		{
			order:       "4",
			description: "fail_username_no_match",
			url:         "https://github.com/f4qr4",
		},
		{
			order:       "5",
			description: "fail_username_no_match",
			url:         "https://instagram.com/r2f2aefa46+_+",
		},
		{
			order:       "6",
			description: "fail_url_not_found",
			url:         "https://facebook/.com/1223e13r",
		},
	}

	for _, url := range urls {
		t.Run(url.order, func(t *testing.T) {
			response, err := urlChecker(url.url)
			if err == schema.ErrSource || err == schema.ErrUsernameNoMatch {
				t.Logf("success, but fail: %s", url.description)
				return
			}

			t.Logf("\ndescription: %s\nresponse: %s\n", url.description, response)
		})
	}
}

func TestCheckUrlUsername(t *testing.T) {
	var urlPort = []struct {
		order       string
		description string
		url         string
	}{
		{
			order:       "1",
			description: "success",
			url:         "/api/check/github/izzxx",
		},
		{
			order:       "2",
			description: "success",
			url:         "/api/check/instagram/jokowi",
		},
		{
			order:       "3",
			description: "success",
			url:         "/api/check/facebook/jokowi",
		},
		{
			order:       "4",
			description: "fail_username_no_match",
			url:         "/api/check/github/.....",
		},
		{
			order:       "5",
			description: "fail_url_not_found",
			url:         "http://localhost:9000/api/url/not/found/check/instagram/jokowi",
		},
		{
			order:       "6",
			description: "fail_souce_discord_not_avaible",
			url:         "/api/check/discord/username",
		},
	}

	for _, url := range urlPort {
		t.Run(url.order, func(t *testing.T) {
			response, err := CheckUrlUsername(url.url)
			if err == schema.ErrUrlFormat || err == schema.ErrSourceNotFound || err == schema.ErrSource || err == schema.ErrUsernameNoMatch {
				t.Logf("success, but fail: %s", url.description)
				return
			}

			t.Logf("\ndescription: %s\nresponse: %s\n", url.description, response)
		})
	}
}
