package lisrp

import (
	"fmt"
)

func MakeDefaultEnv() *Env {
	return &Env{
		Parent: nil,
		Bindings: map[string]Value{
			"+": &PrimitiveFunction{
				Name: "+",
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
		MacroBindings: map[string]Macro{
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
			"lambda": &PrimitiveMacro{
				Name: "lambda",
				ExpandFunc: func(e *SExpression, env *Env) interface{} {
					id_exprs := e.SubExpressions[1].(*SExpression).SubExpressions
					arg_names := make([]string, len(id_exprs))
					for i, _ := range id_exprs {
						arg_names[i] = id_exprs[i].(*Symbol).Id
					}
					return &Function{
						Closure:  env,
						Name:     "<lambda>",
						ArgNames: arg_names,
						Body:     e.SubExpressions[2],
					}
				},
			},
		},
	}
}
