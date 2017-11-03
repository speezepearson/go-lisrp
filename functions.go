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

func (f *Function) Call(env *Env, args []Value) (Value, *LisrpError) {
	if len(f.ArgNames) != len(args) {
		return nil, &LisrpError{fmt.Sprintf("function %s expected %d args, got %d", f.Name, len(f.ArgNames), len(args))}
	}
	new_env := Env{Bindings: map[Symbol]Value{}, Parent: env}
	for i, _ := range args {
		new_env.Bindings[*f.ArgNames[i]] = args[i]
	}
	return f.Body.Eval(&new_env)
}
