package lisrp

import (
	"errors"
	"fmt"
	"strconv"
)

func ParseExpressionWithLeftovers(tokens []Token) (Expression, []Token, error) {
	if len(tokens) == 0 {
		return nil, tokens, errors.New("no tokens to read")
	}
	switch tokens[0].TokenType {
	case TTLParen:
		var subexprs []Expression
		tokens = tokens[1:]
		for len(tokens) > 0 && tokens[0].TokenType != TTRParen {
			var subexpr Expression
			var err error
			subexpr, tokens, err = ParseExpressionWithLeftovers(tokens)
			if err != nil {
				return nil, tokens, err
			}
			subexprs = append(subexprs, subexpr)
		}
		if len(tokens) == 0 {
			return nil, tokens, errors.New("unclosed parenthesis")
		}
		return &SExpression{subexprs}, tokens[1:], nil
	case TTInt:
		n, err := strconv.Atoi(tokens[0].Source)
		if err != nil {
			panic(fmt.Sprintf("expected to be able to parse '%s' as an int", tokens[0].Source))
		}
		return &Integer{n}, tokens[1:], nil
	case TTSymbol:
		return &Symbol{tokens[0].Source}, tokens[1:], nil
	case TTRParen:
		return nil, tokens, errors.New("unexpected RParen")
	}
	panic(fmt.Sprintf("unknown token type %v", tokens[0].TokenType))
}
func ParseTokens(tokens []Token) (Expression, error) {
	result, leftover_tokens, err := ParseExpressionWithLeftovers(tokens)
	if err != nil {
		return nil, err
	}
	if len(leftover_tokens) != 0 {
		return nil, errors.New(fmt.Sprintf("expected EOF; got %v", leftover_tokens))
	}
	return result, nil
}

func Parse(src string) (Expression, error) {
	tokens, err := Tokenize(src)
	if err != nil {
		return nil, err
	}
	result, err := ParseTokens(tokens)
	if err != nil {
		return nil, err
	}
	return result, nil
}
