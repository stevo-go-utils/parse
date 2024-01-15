package parse

import (
	"regexp"
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

func parseStartTag(token html.Token, tagName string, attrs []html.Attribute) bool {
	if tagName != "" && tagName != token.Data {
		return false
	}
	if attrs != nil && !HasAttributes(token, attrs) {
		return false
	}
	return true
}

func parseTextTag(token html.Token, innerHtml string, innerHtmlRegex *regexp.Regexp) bool {
	if innerHtml != "" && token.Data != innerHtml {
		return false
	}
	if innerHtmlRegex != nil && !innerHtmlRegex.MatchString(token.Data) {
		return false
	}
	return true
}
