package lisrp_syntax

import (
	"errors"
	"fmt"
	tok "lisrp/tokenize"
	"strconv"
)

type Symbol string
type Int int

type ExprType string

const (
	Symbol ExprType = "Symbol"
	Int             = "Int"
	SExpr           = "SExpr"
)

type Expr struct {
	ExprType
	SymbolId  string
	NumberVal int
	SExpr     []Expr
}

func ParseExprWithLeftovers(tokens []tok.Token) (Expr, []tok.Token, error) {
	if len(tokens) == 0 {
		return nil, tokens, errors.New("no tokens to read")
	}
	switch tokens[0].TokenType {
	case LParen:
		subexprs := []Expr{}
		tokens = tokens[1:]
		for tokens[0].TokenType != RParen {
			subexpr, tokens, err := ParseExprWithLeftovers(tokens)
			if err != nil {
				return nil, tokens, err
			}
			subexprs = append(subexprs, subexpr)
		}
		return Expr{ExprType: SExpr, SExpr: subexprs}
	case tokens[0].TokenType == Number:
		n, err := strconv.ParseInt(tokens[0].Source)
		return Expr{ExprType: Number, Number: n}
	case Symbol:
		n, err := strconv.ParseInt(tokens[0].Source)
		return Expr{ExprType: Number, Number: n}
	case RParen:
		return nil, tokens[1:], nil
	}
}
func ParseExpr(src string) (Expr, error) {
	result, leftover_tokens, err := ParseExprWithLeftovers(Tokenize(src))
	if err != nil {
		return nil, err
	}
	if len(leftover_tokens) != 0 {
		return nil, errors.new("expected EOF; got {{TODO}}")
	}
	return result, nil
}
