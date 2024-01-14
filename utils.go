package parse

import (
	"slices"

	"golang.org/x/net/html"
)

func HasAttributes(token html.Token, attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if !slices.Contains(token.Attr, attr) {
			return false
		}
	}
	return true
}
