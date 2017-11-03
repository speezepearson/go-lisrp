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
			// "let": &PrimitiveMacro{
			// 	Name: "let",
			// 	ExpandFunc: func(e *SExpression, env *Env) interface{} {
			// 		bindings := map[string]Value{}
			// 		for i, pair := range e.SubExpressions[1].(*SExpression).SubExpressions {
			// 			id := pair.(*SExpression).SubExpressions[0].(*Symbol).Id
			// 			bound_expr := pair.(*SExpression).SubExpressions[1]
			// 			value, lerr := bound_expr.Eval(env)
			// 			if lerr != nil {
			// 				return nil
			// 			}
			// 			bindings[id] = value
			// 		}
			// 		return e.SubExpressions[2].Eval(&Env{Parent: env, Bindings: bindings})
			// 	},
			// },
			Symbol{"lambda"}: &PrimitiveMacro{
				Name: &Symbol{"lambda"},
				ExpandFunc: func(e *SExpression, env *Env) interface{} {
					id_exprs := e.SubExpressions[1].(*SExpression).SubExpressions
					arg_names := make([]*Symbol, len(id_exprs))
					for i, _ := range id_exprs {
						arg_names[i] = id_exprs[i].(*Symbol)
					}
					return &Function{
						Closure:  env,
						Name:     &Symbol{"<lambda>"},
						ArgNames: arg_names,
						Body:     e.SubExpressions[2],
					}
				},
			},
		},
	}
}
