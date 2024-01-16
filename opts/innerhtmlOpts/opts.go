package innerhtmlOpts

import (
	"regexp"

	"golang.org/x/net/html"
)

type Opts struct {
	TagName        string
	InnerHtml      string
	InnerHtmlRegex *regexp.Regexp
	Attrs          []html.Attribute
}

type Opt func(*Opts)

func DefaultOpts() *Opts {
	return &Opts{
		TagName: "",
		Attrs:   nil,
	}
}

func TagNameOpt(tagName string) Opt {
	return func(opts *Opts) {
		opts.TagName = tagName
	}
}

func AttrsOpt(attrs []html.Attribute) Opt {
	return func(opts *Opts) {
		opts.Attrs = attrs
	}
}

func InnerHtmlOpt(innerHtml string) Opt {
	return func(opts *Opts) {
		opts.InnerHtml = innerHtml
	}
}

func InnerHtmlRegexOpt(innerHtmlRegex *regexp.Regexp) Opt {
	return func(opts *Opts) {
		opts.InnerHtmlRegex = innerHtmlRegex
	}
}
