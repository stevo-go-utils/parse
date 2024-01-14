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

type AttrValOptFunc func(*AttrValOpts)

func DefaultAttrValOpts() *AttrValOpts {
	return &AttrValOpts{
		tagName: "",
		attrs:   nil,
	}
}

func AttrValTagNameOpt(tagName string) AttrValOptFunc {
	return func(opts *AttrValOpts) {
		opts.tagName = tagName
	}
}

func AttrValAttrsOpt(attrs []html.Attribute) AttrValOptFunc {
	return func(opts *AttrValOpts) {
		opts.attrs = attrs
	}
}

func InnerHtmlOpt(innerHtml string) AttrValOptFunc {
	return func(opts *AttrValOpts) {
		opts.innerHtml = innerHtml
	}
}

func InnerHtmlRegexOpt(innerHtmlRegex *regexp.Regexp) AttrValOptFunc {
	return func(opts *AttrValOpts) {
		opts.innerHtmlRegex = innerHtmlRegex
	}
}

func AttrVal(body string, attrKey string, opts ...AttrValOptFunc) (val string, err error) {
	defaultOpts := DefaultAttrValOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	tkn := html.NewTokenizer(strings.NewReader(body))
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			err = fmt.Errorf("failed to find matching token with attr: %s", attrKey)
			return
		case html.StartTagToken:
			token := tkn.Token()
			if !attrValTokenCheck(token, defaultOpts) {
				continue
			}
			for _, attr := range token.Attr {
				if attr.Key == attrKey {
					val = attr.Val
					return
				}
			}
		}
	}
}

func AttrVals(body string, attrKey string, opts ...AttrValOptFunc) (vals []string) {
	defaultOpts := DefaultAttrValOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	tkn := html.NewTokenizer(strings.NewReader(body))
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			token := tkn.Token()
			if !attrValTokenCheck(token, defaultOpts) {
				continue
			}
			for _, attr := range token.Attr {
				if attr.Key == attrKey {
					vals = append(vals, attr.Val)
				}
			}
		}
	}
}

func attrValTokenCheck(token html.Token, opts *AttrValOpts) bool {
	if opts.tagName != "" && opts.tagName != token.Data {
		return false
	}
	if opts.attrs != nil && !HasAttributes(token, opts.attrs) {
		return false
	}
	if opts.innerHtml != "" && token.Data != opts.innerHtml {
		return false
	}
	if opts.innerHtmlRegex != nil && !opts.innerHtmlRegex.MatchString(token.Data) {
		return false
	}
	return true
}
