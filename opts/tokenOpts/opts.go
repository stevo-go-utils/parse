package tokenOpts

import "golang.org/x/net/html"

type Opts struct {
	TokenType html.TokenType
}

type Opt func(*Opts)

func DefaultOpts() *Opts {
	return &Opts{
		TokenType: html.TokenType(^uint32(0)),
	}
}

func TokenTypeOpt(tokenType html.TokenType) Opt {
	return func(opts *Opts) {
		opts.TokenType = tokenType
	}
}
