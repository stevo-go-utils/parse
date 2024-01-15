package parse

import (
	"fmt"
	"slices"
	"strings"

	"github.com/stevo-go-utils/parse/opts/attributeOpts"
	"golang.org/x/net/html"
)

func AttrVal(body string, attrKey string, opts ...attributeOpts.Opt) (val string, err error) {
	defaultOpts := attributeOpts.DefaultOpts()
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

func AttrVals(body string, attrKey string, opts ...attributeOpts.Opt) (vals []string) {
	defaultOpts := attributeOpts.DefaultOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	return parseAttrVal(html.NewTokenizer(strings.NewReader(body)), attrKey, defaultOpts)
}

func parseAttrVal(tkn *html.Tokenizer, attrKey string, opts *attributeOpts.Opts) (vals []string) {
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
			if !parseStartTag(token, opts.TagName, opts.Attrs) {
				continue
			}
			if opts.InnerHtml != "" || opts.InnerHtmlRegex != nil {
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
			if !parseTextTag(token, opts.InnerHtml, opts.InnerHtmlRegex) {
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

func AttrsVal(body string, attrKeys []string, opts ...attributeOpts.Opt) (val map[string]string, err error) {
	defaultOpts := attributeOpts.DefaultOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	vals := parseAttrsVal(html.NewTokenizer(strings.NewReader(body)), attrKeys, defaultOpts)
	if len(vals) == 0 {
		err = fmt.Errorf("no value found for attributes %s", attrKeys)
		return
	}
	val = vals[0]
	return
}

func AttrsVals(body string, attrKeys []string, opts ...attributeOpts.Opt) (vals []map[string]string) {
	defaultOpts := attributeOpts.DefaultOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	return parseAttrsVal(html.NewTokenizer(strings.NewReader(body)), attrKeys, defaultOpts)
}

func parseAttrsVal(tkn *html.Tokenizer, attrKeys []string, opts *attributeOpts.Opts) (vals []map[string]string) {
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
			if !parseStartTag(token, opts.TagName, opts.Attrs) {
				continue
			}
			if opts.InnerHtml != "" || opts.InnerHtmlRegex != nil {
				checkText = true
				prevToken = &token
				continue
			}
			if !HasAttributeKeys(token, attrKeys) {
				continue
			}
			val := make(map[string]string)
			for _, attr := range token.Attr {
				if slices.Contains(attrKeys, attr.Key) {
					val[attr.Key] = attr.Val
				}
			}
			vals = append(vals, val)
		case html.TextToken:
			if !checkText || prevToken == nil {
				continue
			}
			token := tkn.Token()
			if !parseTextTag(token, opts.InnerHtml, opts.InnerHtmlRegex) {
				continue
			}
			if !HasAttributeKeys(token, attrKeys) {
				continue
			}
			val := make(map[string]string)
			for _, attr := range token.Attr {
				if slices.Contains(attrKeys, attr.Key) {
					val[attr.Key] = attr.Val
				}
			}
			vals = append(vals, val)
		}
		checkText = false
		prevToken = nil
	}
}
