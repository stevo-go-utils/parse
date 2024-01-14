package parse

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type InnerHtmlOpts struct {
	tagName string
	attrs   []html.Attribute
}

type InnerHtmlOptFunc func(*InnerHtmlOpts)

func DefaultInnerHtmlOpts() *InnerHtmlOpts {
	return &InnerHtmlOpts{
		tagName: "",
		attrs:   nil,
	}
}

func InnerHtmlTagNameOpt(tagName string) InnerHtmlOptFunc {
	return func(opts *InnerHtmlOpts) {
		opts.tagName = tagName
	}
}

func InnerHtmlAttrsOpt(attrs []html.Attribute) InnerHtmlOptFunc {
	return func(opts *InnerHtmlOpts) {
		opts.attrs = attrs
	}
}

func InnerHtml(body string, attrKey string, opts ...InnerHtmlOptFunc) (innerHtml string, err error) {
	defaultOpts := DefaultInnerHtmlOpts()
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
			if defaultOpts.tagName != "" && defaultOpts.tagName != token.Data {
				continue
			}
			if defaultOpts.attrs != nil {
				var found bool
				for _, attr := range token.Attr {
					if attr.Key == attrKey {
						found = true
						break
					}
				}
				if !found {
					continue
				}
			}
			innerHtml = token.Data
			return
		}
	}
}
