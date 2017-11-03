package lisrp

import (
	"fmt"
)

type Callable interface {
	Call(*Env, []Value) (Value, *LisrpError)
}

type Function struct {
	Closure  *Env
	Name     *Symbol
	ArgNames []*Symbol
	Body     Expression
}

func (f *Function) String() string {
	return fmt.Sprintf("<function %s>", f.Name)
}

func (f *Function) Call(_ *Env, args []Value) (Value, *LisrpError) {
	if len(f.ArgNames) != len(args) {
		return nil, &LisrpError{fmt.Sprintf("function %s expected %d args, got %d", f.Name, len(f.ArgNames), len(args))}
	}
	new_env := Env{Bindings: map[Symbol]Value{}, Parent: f.Closure}
	for i, _ := range args {
		new_env.Bindings[*f.ArgNames[i]] = args[i]
	}
	return f.Body.Eval(&new_env)
}

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
