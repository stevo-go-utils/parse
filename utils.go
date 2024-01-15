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

func HasAttributeKeys(token html.Token, attrKeys []string) bool {
	var tokenKeys []string
	for _, attr := range token.Attr {
		tokenKeys = append(tokenKeys, attr.Key)
	}
	for _, attrKey := range attrKeys {
		if !slices.Contains(tokenKeys, attrKey) {
			return false
		}
	}
	return true
}
