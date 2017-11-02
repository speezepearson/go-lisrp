package lisrp

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
	TTSymbol TokenType = "Symbol"
	TTInt              = "Int"
	TTLParen           = "("
	TTRParen           = ")"
)

type Token struct {
	TokenType
	Source string
}

var QuoteRegexp = regexp.MustCompile(`^'`)
var LParenRegexp = regexp.MustCompile(`^[({\[]`)
var RParenRegexp = regexp.MustCompile(`^[)}\]]`)
var NumberRegexp = regexp.MustCompile(`^[+-]?[0-9]+`)
var SymbolRegexp = regexp.MustCompile(`^[a-zA-Z0-9~!@#$%^&*_+=|:;'",<.>/?-]+`)

func Tokenize(src string) ([]Token, error) {
	var result []Token
	src = strings.TrimSpace(src)
	for len(src) > 0 {
		var match string
		if matches := LParenRegexp.FindStringSubmatch(src); matches != nil {
			match = matches[0]
			result = append(result, Token{TokenType: TTLParen, Source: match})
		} else if matches := RParenRegexp.FindStringSubmatch(src); matches != nil {
			match = matches[0]
			result = append(result, Token{TokenType: TTRParen, Source: match})
		} else if matches := NumberRegexp.FindStringSubmatch(src); matches != nil {
			match = matches[0]
			result = append(result, Token{TokenType: TTInt, Source: match})
		} else if matches := SymbolRegexp.FindStringSubmatch(src); matches != nil {
			match = matches[0]
			result = append(result, Token{TokenType: TTSymbol, Source: match})
		} else {
			return nil, MakeNonTokenError(src)
		}
		src = strings.TrimSpace(strings.TrimPrefix(src, match))
	}
	return result, nil
}
