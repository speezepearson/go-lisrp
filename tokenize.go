package lisrp_syntax

import (
	"errors"
	"regexp"
	"strings"
)

func MakeNonTokenError(src string) error {
	return errors.New("no token can match: " + src)
}

type TokenType string

const (
	Symbol TokenType = "Symbol"
	Int              = "Int"
	LParen           = "("
	RParen           = ")"
)

type Token struct {
	TokenType
	Source string
}

const (
	LParenRegexp = regexp.Compile(`[({\[]`)
	RParenRegexp = regexp.Compile(`[)}\]]`)
	NumberRegexp = regexp.Compile(`[+-]?[0-9]+`)
	SymbolRegexp = regexp.Compile(`[a-zA-Z0-9~!@#$%^&*()_+=|:;'",<.>/?-]+`)
)

func Tokenize(src string) ([]Token, error) {
	var result []Token
	src = strings.TrimSpace(src)
	for len(src) > 0 {
		if matches := regexp.FindStringMatch(LParenRegexp, src); matches != nil {
			append(result, Token{TokenType: LParen, Source: matches[0]})
		} else if matches := regexp.FindStringMatch(RParenRegexp, src); matches != nil {
			append(result, Token{TokenType: RParen, Source: matches[0]})
		} else if matches := regexp.FindStringMatch(NumberRegexp, src); matches != nil {
			append(result, Token{TokenType: Number, Source: matches[0]})
		} else if matches := regexp.FindStringMatch(SymbolRegexp, src); matches != nil {
			append(result, Token{TokenType: Symbol, Source: matches[0]})
		} else {
			return nil, MakeNonTokenError(src)
		}
		src = strings.TrimSpace(src)
	}
	return result, nil
}
