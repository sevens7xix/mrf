package mrf

import "strings"

func stringFormatter(rawString []string) string {

	if len(rawString) > 1 {
		return strings.Join(rawString, "+")
	} else {
		return rawString[0]
	}
}
