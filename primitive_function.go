package lisrp

import (
	"fmt"
)

type PrimitiveFunction struct {
	Name string
	Code func(*Env, []Value) (Value, *LisrpError)
}

func (f *PrimitiveFunction) String() string {
	return fmt.Sprintf("<function %s>", f.Name)
}

func (f PrimitiveFunction) Call(env *Env, args []Value) (Value, *LisrpError) {
	return f.Code(env, args)
}
