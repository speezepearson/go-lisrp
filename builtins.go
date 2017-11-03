package lisrp

import (
	"fmt"
)

func MakeDefaultEnv() *Env {
	return &Env{
		Parent: nil,
		Bindings: map[string]Value{
			"+": PrimitiveFunction{
				Name: "+",
				Code: func(env *Env, args []Value) (Value, *LisrpError) {
					result := 0
					for _, arg := range args {
						n, ok := (arg).(Integer)
						if !ok {
							return nil, &LisrpError{fmt.Sprintf("can only add ints, not %v", arg)}
						}
						result += n.Value
					}
					return &Integer{result}, nil
				},
			},

			"lambda": PrimitiveFunction{
				Name: "lambda",
				Code: func(env *Env, args []Value) (Value, *LisrpError) {
					panic("macros not implemented")
				},
			},
		},
	}
}
