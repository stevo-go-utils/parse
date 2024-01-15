package parse

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func AttrsVal(body string, attrKeys []string, opts ...AttrValOptFunc) (val map[string]string, err error) {
	defaultOpts := DefaultAttrValOpts()
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

func AttrsVals(body string, attrKeys []string, opts ...AttrValOptFunc) (vals []map[string]string) {
	defaultOpts := DefaultAttrValOpts()
	for _, opt := range opts {
		opt(defaultOpts)
	}
	return parseAttrsVal(html.NewTokenizer(strings.NewReader(body)), attrKeys, defaultOpts)
}

func parseAttrsVal(tkn *html.Tokenizer, attrKeys []string, opts *AttrValOpts) (vals []map[string]string) {
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
			if !attrsValTokenCheck(token, opts) {
				continue
			}
			if opts.innerHtml != "" || opts.innerHtmlRegex != nil {
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
			if !attrsValTextCheck(token, opts) {
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

func attrsValTokenCheck(token html.Token, opts *AttrValOpts) bool {
	if opts.tagName != "" && opts.tagName != token.Data {
		return false
	}
	if opts.attrs != nil && !HasAttributes(token, opts.attrs) {
		return false
	}
	return true
}

func attrsValTextCheck(token html.Token, opts *AttrValOpts) bool {
	if opts.innerHtml != "" && token.Data != opts.innerHtml {
		return false
	}
	if opts.innerHtmlRegex != nil && !opts.innerHtmlRegex.MatchString(token.Data) {
		return false
	}
	return true
}
