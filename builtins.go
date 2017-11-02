package lisrp

import (
	"fmt"
)

func MakeDefaultEnv() *Env {
	return &Env{
		Parent: nil,
		Bindings: map[string]*Value{
			"+": &Value{ValueType: VTPrimitiveFunction, PrimitiveFunction: PrimitiveFunction{
				Name: "+",
				Call: func(args []*Value, env *Env) (*Value, *LisrpError) {
					result := 0
					for _, arg := range args {
						if arg.ValueType != VTInt {
							return nil, &LisrpError{fmt.Sprintf("can only add ints, not %v", arg)}
						}
						result += arg.Int
					}
					return &Value{ValueType: VTInt, Int: result}, nil
				},
			}},

			"-": &Value{ValueType: VTPrimitiveFunction, PrimitiveFunction: PrimitiveFunction{
				Name: "-",
				Call: func(args []*Value, env *Env) (*Value, *LisrpError) {
					result := 0
					if len(args) > 1 && args[0].ValueType == VTInt {
						result = args[0].Int
						args = args[1:]
					}
					for _, arg := range args {
						if arg.ValueType != VTInt {
							return nil, &LisrpError{fmt.Sprintf("can only multiply ints, not %v", arg)}
						}
						result -= arg.Int
					}
					return &Value{ValueType: VTInt, Int: result}, nil
				},
			}},

			"*": &Value{ValueType: VTPrimitiveFunction, PrimitiveFunction: PrimitiveFunction{
				Name: "*",
				Call: func(args []*Value, env *Env) (*Value, *LisrpError) {
					result := 1
					for _, arg := range args {
						if arg.ValueType != VTInt {
							return nil, &LisrpError{fmt.Sprintf("can only multiply ints, not %v", arg)}
						}
						result *= arg.Int
					}
					return &Value{ValueType: VTInt, Int: result}, nil
				},
			}},
		},
	}
}
