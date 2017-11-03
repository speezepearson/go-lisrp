package lisrp

import (
	"fmt"
)

type SpecialForm struct {
	Name *Symbol
	Eval func(*Env, []Expression) (Value, *LisrpError)
}

var SpecialForms = map[Symbol]SpecialForm{
	Symbol{"let1"}: SpecialForm{
		Name: &Symbol{"let1"},
		Eval: func(env *Env, args []Expression) (Value, *LisrpError) {
			id := args[0].(*Symbol)
			bound_value, lerr := args[1].Eval(env)
			if lerr != nil {
				return nil, nil
			}
			return args[2].Eval(&Env{Parent: env, Bindings: map[Symbol]Value{*id: bound_value}})
		},
	},
	Symbol{"unary-function"}: SpecialForm{
		Name: &Symbol{"unary-function"},
		Eval: func(env *Env, args []Expression) (Value, *LisrpError) {
			fname := args[0].(*Symbol)
			argname := args[1].(*Symbol)
			return &Function{
				Closure:  env,
				Name:     fname,
				ArgNames: []*Symbol{argname},
				Body:     args[2],
			}, nil
		},
	},
	Symbol{"define-symbol"}: SpecialForm{
		Name: &Symbol{"define-symbol"},
		Eval: func(env *Env, args []Expression) (Value, *LisrpError) {
			if len(args) != 2 {
				return nil, &LisrpError{"usage: (define ID EXPR)"}
			}
			id, id_is_symbol := args[0].(*Symbol)
			if !id_is_symbol {
				return nil, &LisrpError{fmt.Sprintf("%v is not a symbol", args[0])}
			}
			_, already_present := env.Bindings[*id]
			if already_present {
				return nil, &LisrpError{fmt.Sprintf("%v is already defined", id)}
			}
			val, lerr := args[1].Eval(env)
			if lerr != nil {
				return nil, lerr
			}
			env.Bindings[*id] = val
			return Void, nil
		},
	},
	Symbol{"set!"}: SpecialForm{
		Name: &Symbol{"set!"},
		Eval: func(env *Env, args []Expression) (Value, *LisrpError) {
			if len(args) != 2 {
				return nil, &LisrpError{"usage: (set! ID EXPR)"}
			}
			id, id_is_symbol := args[0].(*Symbol)
			if !id_is_symbol {
				return nil, &LisrpError{fmt.Sprintf("%v not a symbol", args[0])}
			}
			for env != nil {
				_, present := env.Bindings[*id]
				if present {
					val, lerr := args[1].Eval(env)
					if lerr != nil {
						return nil, lerr
					}
					env.Bindings[*id] = val
					return Void, nil
				}
				env = env.Parent
			}
			return nil, &LisrpError{fmt.Sprintf("cannot set undefined variable %v", id)}
		},
	},
	Symbol{"begin"}: SpecialForm{
		Name: &Symbol{"begin"},
		Eval: func(env *Env, args []Expression) (Value, *LisrpError) {
			if len(args) == 0 {
				return nil, &LisrpError{"`begin` form must not be empty"}
			}
			var result Value
			for _, arg := range args {
				var lerr *LisrpError
				result, lerr = arg.Eval(env)
				if lerr != nil {
					return nil, lerr
				}
			}
			return result, nil
		},
	},
	Symbol{"begin0"}: SpecialForm{
		Name: &Symbol{"begin0"},
		Eval: func(env *Env, args []Expression) (Value, *LisrpError) {
			if len(args) == 0 {
				return nil, &LisrpError{"`begin0` form must not be empty"}
			}
			result, lerr := args[0].Eval(env)
			if lerr != nil {
				return nil, lerr
			}
			for _, arg := range args[1:] {
				_, lerr = arg.Eval(env)
				if lerr != nil {
					return nil, lerr
				}
			}
			return result, nil
		},
	},
}

func (sf *SpecialForm) String() string {
	return fmt.Sprintf("<builtin %s>", sf.Name.Id)
}
