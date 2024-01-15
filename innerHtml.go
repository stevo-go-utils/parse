package parse

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type InnerHtmlOpts struct {
	tagName        string
	innerHtml      string
	innerHtmlRegex *regexp.Regexp
	attrs          []html.Attribute
}

type InnerHtmlOpt func(*InnerHtmlOpts)

func DefaultInnerHtmlOpts() *InnerHtmlOpts {
	return &InnerHtmlOpts{
		tagName: "",
		attrs:   nil,
	}
}

func TagNameInnerHtmlOpt(tagName string) InnerHtmlOpt {
	return func(opts *InnerHtmlOpts) {
		opts.tagName = tagName
	}
}

func AttrsInnerHtmlOpt(attrs []html.Attribute) InnerHtmlOpt {
	return func(opts *InnerHtmlOpts) {
		opts.attrs = attrs
	}
}

func InnerHtmlInnerHtmlOpt(innerHtml string) InnerHtmlOpt {
	return func(opts *InnerHtmlOpts) {
		opts.innerHtml = innerHtml
	}
}

func InnerHtmlRegexInnerHtmlOpt(innerHtmlRegex *regexp.Regexp) InnerHtmlOpt {
	return func(opts *InnerHtmlOpts) {
		opts.innerHtmlRegex = innerHtmlRegex
	}
}

func InnerHtml(body string, opts ...InnerHtmlOpt) (innerHtml string, err error) {
	var checkText bool = false
	defaultOpts := DefaultInnerHtmlOpts()
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
			if parseStartTag(token, defaultOpts.tagName, defaultOpts.attrs) {
				checkText = true
				continue
			}
		case html.TextToken:
			if !checkText {
				continue
			}
			token := tkn.Token()
			if !parseTextTag(token, defaultOpts.innerHtml, defaultOpts.innerHtmlRegex) {
				continue
			}
			innerHtml = token.Data
			return
		}
		checkText = false
	}
}
