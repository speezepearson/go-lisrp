package lisrp

import (
	"fmt"
)

type Function struct {
	Closure  *Env
	Name     string
	ArgNames []string
	Body     Expression
}

func (f *Function) String() string {
	return fmt.Sprintf("<function %s>", f.Name)
}

func (f Function) Call(env *Env, args []Value) (Value, *LisrpError) {
	if len(f.ArgNames) != len(args) {
		return nil, &LisrpError{fmt.Sprintf("function %s expected %d args, got %d", f.Name, len(f.ArgNames), len(args))}
	}
	new_env := Env{Parent: env}
	for i, _ := range args {
		new_env.Bindings[f.ArgNames[i]] = args[i]
	}
	return f.Body.Eval(&new_env)
}
