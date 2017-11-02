package lisrp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Symbol string
type Int int

type ExprType string

const (
	ETSymbol ExprType = "Symbol"
	ETInt             = "Int"
	ETSExpr           = "SExpr"
)

type Expr struct {
	ExprType
	Symbol   string
	Int      int
	SubExprs []Expr
}

func ParseExprWithLeftovers(tokens []Token) (Expr, []Token, error) {
	if len(tokens) == 0 {
		return Expr{}, tokens, errors.New("no tokens to read")
	}
	switch tokens[0].TokenType {
	case TTLParen:
		subexprs := []Expr{}
		tokens = tokens[1:]
		for len(tokens) > 0 && tokens[0].TokenType != TTRParen {
			var subexpr Expr
			var err error
			subexpr, tokens, err = ParseExprWithLeftovers(tokens)
			if err != nil {
				return Expr{}, tokens, err
			}
			subexprs = append(subexprs, subexpr)
		}
		if len(tokens) == 0 {
			return Expr{}, tokens, errors.New("unclosed parenthesis")
		}
		return Expr{ExprType: ETSExpr, SubExprs: subexprs}, tokens[1:], nil
	case TTInt:
		n, err := strconv.Atoi(tokens[0].Source)
		if err != nil {
			panic(fmt.Sprintf("expected to be able to parse '%s' as an int", tokens[0].Source))
		}
		return Expr{ExprType: ETInt, Int: n}, tokens[1:], nil
	case TTSymbol:
		return Expr{ExprType: ETSymbol, Symbol: tokens[0].Source}, tokens[1:], nil
	case TTRParen:
		return Expr{}, tokens, errors.New("unexpected RParen")
	}
	panic(fmt.Sprintf("unknown token type %v", tokens[0].TokenType))
}
func ParseTokens(tokens []Token) (Expr, error) {
	result, leftover_tokens, err := ParseExprWithLeftovers(tokens)
	if err != nil {
		return Expr{}, err
	}
	if len(leftover_tokens) != 0 {
		return Expr{}, errors.New(fmt.Sprintf("expected EOF; got %v", leftover_tokens))
	}
	return result, nil
}

func Parse(src string) (Expr, error) {
	tokens, err := Tokenize(src)
	if err != nil {
		return Expr{}, err
	}
	result, err := ParseTokens(tokens)
	if err != nil {
		return Expr{}, err
	}
	return result, nil
}

func (e Expr) String() string {
	switch e.ExprType {
	case ETSymbol:
		return e.Symbol
	case ETInt:
		return strconv.Itoa(e.Int)
	case ETSExpr:
		words := make([]string, len(e.SubExprs))
		for i, e := range e.SubExprs {
			words[i] = e.String()
		}
		return fmt.Sprintf("(%s)", strings.Join(words, " "))
	}
	panic(fmt.Sprintf("unknown expression type %v", e.ExprType))
}
