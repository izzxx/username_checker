package utility

import "testing"

func TestValidate(t *testing.T) {
	var urls = []struct {
		order       string
		description string
		url         string
	}{
		{
			order:       "1",
			description: "success",
			url:         "/api/check/instagram/{username}",
		},
		{
			order:       "2",
			description: "fail",
			url:         "http://localhost:9000/api/check/what?/twiter/{username}",
		},
		{
			order:       "3",
			description: "success",
			url:         "/api/check/facebook/{username}",
		},
		{
			order:       "4",
			description: "fail",
			url:         "http://localhost:9000/api/check/twiter/{username}/{username}",
		},
	}

	for _, url := range urls {
		t.Run(url.order, func(t *testing.T) {
			detail, username, ok := ValidateUrl(url.url)
			if !ok || url.description == "fail" {
				t.Log("success, but fail")
				return
			}

			t.Logf("\ndetail: %s\nusername: %s\n", detail, username)
		})
	}
}
