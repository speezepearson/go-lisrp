package lisrp

import (
	"fmt"
)

type PrimitiveFunction struct {
	Name *Symbol
	Code func(*Env, []Value) (Value, *LisrpError)
}

func (f *PrimitiveFunction) String() string {
	return fmt.Sprintf("<function %s>", f.Name.Id)
}

func (f *PrimitiveFunction) Call(env *Env, args []Value) (Value, *LisrpError) {
	return f.Code(env, args)
}
