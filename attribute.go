package parse

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type AttrValOpts struct {
	tagName        string
	innerHtml      string
	innerHtmlRegex *regexp.Regexp
	attrs          []html.Attribute
}

type AttrValOpt func(*AttrValOpts)

func DefaultAttrValOpts() *AttrValOpts {
	return &AttrValOpts{
		tagName: "",
		attrs:   nil,
	}
}

func TagNameAttrValOpt(tagName string) AttrValOpt {
	return func(opts *AttrValOpts) {
		opts.tagName = tagName
	}
}

func AttrsAttrValOpt(attrs []html.Attribute) AttrValOpt {
	return func(opts *AttrValOpts) {
		opts.attrs = attrs
	}
}

func InnerHtmlAttrValOpt(innerHtml string) AttrValOpt {
	return func(opts *AttrValOpts) {
		opts.innerHtml = innerHtml
	}
}

func InnerHtmlRegexAttrValOpt(innerHtmlRegex *regexp.Regexp) AttrValOpt {
	return func(opts *AttrValOpts) {
		opts.innerHtmlRegex = innerHtmlRegex
	}
}

func AttrVal(body string, attrKey string, opts ...AttrValOpt) (val string, err error) {
	defaultOpts := DefaultAttrValOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	vals := parseAttrVal(html.NewTokenizer(strings.NewReader(body)), attrKey, defaultOpts)
	if len(vals) == 0 {
		return "", fmt.Errorf("no value found for attribute '%s'", attrKey)
	}
	val = vals[0]
	return
}

func AttrVals(body string, attrKey string, opts ...AttrValOpt) (vals []string) {
	defaultOpts := DefaultAttrValOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	return parseAttrVal(html.NewTokenizer(strings.NewReader(body)), attrKey, defaultOpts)
}

func parseAttrVal(tkn *html.Tokenizer, attrKey string, opts *AttrValOpts) (vals []string) {
	var (
		prevToken *html.Token = nil
		checkText bool        = false
	)
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			token := tkn.Token()
			if !parseStartTag(token, opts.tagName, opts.attrs) {
				continue
			}
			if opts.innerHtml != "" || opts.innerHtmlRegex != nil {
				checkText = true
				prevToken = &token
				continue
			}
			for _, attr := range token.Attr {
				if attr.Key == attrKey {
					vals = append(vals, attr.Val)
				}
			}
		case html.TextToken:
			if !checkText || prevToken == nil {
				continue
			}
			token := tkn.Token()
			if !parseTextTag(token, opts.innerHtml, opts.innerHtmlRegex) {
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
