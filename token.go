package parse

import (
	"errors"
	"strings"

	"github.com/stevo-go-utils/parse/opts/tokenOpts"
	"golang.org/x/net/html"
)

func Token(body string, opts ...tokenOpts.Opt) (token html.Token, err error) {
	defaultOpts := tokenOpts.DefaultOpts()
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
		default:

		}
	}
}
