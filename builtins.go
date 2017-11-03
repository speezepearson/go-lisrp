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
	}
}
