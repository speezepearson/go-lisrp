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
	Symbol{"lambda1"}: SpecialForm{
		Name: &Symbol{"lambda1"},
		Eval: func(env *Env, args []Expression) (Value, *LisrpError) {
			arg_id := args[0].(*Symbol)
			return &Function{
				Closure:  env,
				Name:     &Symbol{"<lambda>"},
				ArgNames: []*Symbol{arg_id},
				Body:     args[1],
			}, nil
		},
	},
}

func (sf *SpecialForm) String() string {
	return fmt.Sprintf("<builtin %s>", sf.Name.Id)
}
