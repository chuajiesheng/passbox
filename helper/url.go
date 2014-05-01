package helper

import "strings"

func HandleURL(requestURL string, url string) bool {
	return strings.HasSuffix(requestURL, url)
}
