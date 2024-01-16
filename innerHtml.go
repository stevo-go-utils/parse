package parse

import (
	"errors"
	"strings"

	"github.com/stevo-go-utils/parse/opts/innerhtmlOpts"
	"golang.org/x/net/html"
)

func InnerHtml(body string, opts ...innerhtmlOpts.Opt) (innerHtml string, err error) {
	var checkText bool = false
	defaultOpts := innerhtmlOpts.DefaultOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	tkn := html.NewTokenizer(strings.NewReader(body))
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			err = errors.New("no inner html found")
			return
		case html.StartTagToken:
			token := tkn.Token()
			if parseStartTag(token, defaultOpts.TagName, defaultOpts.Attrs) {
				checkText = true
				continue
			}
		case html.TextToken:
			if !checkText {
				continue
			}
			token := tkn.Token()
			if !parseTextTag(token, defaultOpts.InnerHtml, defaultOpts.InnerHtmlRegex) {
				continue
			}
			innerHtml = token.Data
			return
		}
		checkText = false
	}
}

func InnerHtmls(body string, opts ...innerhtmlOpts.Opt) (innerHtml string, err error) {
	var checkText bool = false
	defaultOpts := innerhtmlOpts.DefaultOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	tkn := html.NewTokenizer(strings.NewReader(body))
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			err = errors.New("no inner html found")
			return
		case html.StartTagToken:
			token := tkn.Token()
			if parseStartTag(token, defaultOpts.TagName, defaultOpts.Attrs) {
				checkText = true
				continue
			}
		case html.TextToken:
			if !checkText {
				continue
			}
			token := tkn.Token()
			if !parseTextTag(token, defaultOpts.InnerHtml, defaultOpts.InnerHtmlRegex) {
				continue
			}
			innerHtml = token.Data
			return
		}
		checkText = false
	}
}
