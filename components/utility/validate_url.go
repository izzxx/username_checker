package utility

import (
	"strings"
)

func ValidateUrl(url string) (string, string, bool) {
	split := strings.Split(url, "/")

	if len(split) > 5 {
		return "", "", false
	}

	// detail is search source
	detail := split[3]

	// username is the username in the source
	username := split[4]

	return detail, username, true
}
