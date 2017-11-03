package lisrp

import (
	"fmt"
)

func MakeDefaultEnv() *Env {
	return &Env{
		Parent: nil,
		Bindings: map[Symbol]Value{
			Symbol{"+"}: &PrimitiveFunction{
				Name: &Symbol{"+"},
				Code: func(env *Env, args []Value) (Value, *LisrpError) {
					result := 0
					for _, arg := range args {
						n, ok := (arg).(*Integer)
						if !ok {
							return nil, &LisrpError{fmt.Sprintf("can only add ints, not %v", arg)}
						}
						result += n.Value
					}
					return &Integer{result}, nil
				},
			},
		},
		MacroBindings: map[Symbol]Macro{
			Symbol{"define"}: &PrimitiveMacro{
				Name: &Symbol{"define"},
				ExpandFunc: func(e *SExpression, env *Env) interface{} {
					var id *Symbol
					var expr Expression
					switch assignee := e.SubExpressions[1].(type) {
					case *Symbol:
						id = assignee
						expr = e.SubExpressions[2]
					case *SExpression:
						id = assignee.SubExpressions[0].(*Symbol)
						expr = &SExpression{[]Expression{
							&Symbol{"unary-function"},
							id,
							assignee.SubExpressions[1].(*Symbol),
							e.SubExpressions[2],
						}}
					})
					return &SExpression{[]Expression{
						&Symbol{"define-symbol"},
						id,
						expr,
					}}
				},
			},
		},
	}
}
