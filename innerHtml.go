package parse

import (
	"errors"
	"strings"

	"golang.org/x/net/html"
)

type InnerHtmlOpts struct {
	tagName string
	attrs   []html.Attribute
}

type InnerHtmlOpt func(*InnerHtmlOpts)

func DefaultInnerHtmlOpts() *InnerHtmlOpts {
	return &InnerHtmlOpts{
		tagName: "",
		attrs:   nil,
	}
}

func InnerHtml(body string, opts ...InnerHtmlOpt) (innerHtml string, err error) {
	defaultOpts := DefaultInnerHtmlOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	tkn := html.NewTokenizer(strings.NewReader(body))
	for {
		for {
			tt := tkn.Next()
			switch tt {
			case html.ErrorToken:
				return
			case html.StartTagToken:
				token := tkn.Token()
				if !attrValTokenCheck(token, opts) {
					continue
				}
				if opts.innerHtml != "" || opts.innerHtmlRegex != nil {
					checkText = true
					prevToken = &token
					continue
				}
			case html.TextToken:
				if !checkText || prevToken == nil {
					continue
				}
				token := tkn.Token()
				if !attrValTextCheck(token, opts) {
					continue
				}
				for _, attr := range prevToken.Attr {
					if attr.Key == attrKey {
						vals = append(vals, attr.Val)
					}
				}
			}
			checkText = false
			prevToken = nil
		}
}

/*
func InnerHtmlValue(input string, tagName string, attributes []Attribute, contains string) ([]string, error) {
	var res []string
	for _, attr := range attributes {
		if attr.Key == "" && attr.Val == "" {
			return res, errors.New("empty attribute check entered")
		}
	}
	tkn := html.NewTokenizer(strings.NewReader(input))
	isToken := false
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			return res, nil
		case html.StartTagToken:
			token := tkn.Token()
			if token.Data == tagName {
				var successfulChecks int
				for _, check := range attributes {
					if checkAttributes(token.Attr, check) {
						successfulChecks++
					}
				}
				if successfulChecks == len(attributes) {
					isToken = true
				}
			}
		case html.TextToken:
			if isToken {
				token := tkn.Token()
				if strings.ContainsAny(token.Data, `!@#$%^&*()_+-={}[]:;"'?/>.<,QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm|\1234567890`+"`") {
					if contains != "" {
						if strings.Contains(token.Data, contains) {
							res = append(res, token.Data)
						}
					} else {
						res = append(res, token.Data)
					}
				}
			}
			isToken = false
		}
	}
}

*/
