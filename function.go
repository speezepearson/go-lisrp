package lisrp

import (
	"fmt"
)

type Callable interface {
	Call(*Env, []LisrpValue) (LisrpValue, *LisrpError)
}

var FunctionT = LisrpType{"function"}

type Function struct {
	Closure  *Env
	Name     *Symbol
	ArgNames []*Symbol
	Body     Expression
}

func (f *Function) Type() LisrpType {
	return FunctionT
}
func (f *Function) String() string {
	return fmt.Sprintf("<function %s>", f.Name)
}

func (f *Function) Call(_ *Env, args []LisrpValue) (LisrpValue, *LisrpError) {
	if len(f.ArgNames) != len(args) {
		return nil, &LisrpError{fmt.Sprintf("function %s expected %d args, got %d", f.Name, len(f.ArgNames), len(args))}
	}
	new_env := Env{Bindings: map[Symbol]LisrpValue{}, Parent: f.Closure}
	for i, _ := range args {
		new_env.Bindings[*f.ArgNames[i]] = args[i]
	}
	return f.Body.Eval(&new_env)
}

type PrimitiveFunction struct {
	Name *Symbol
	Code func(*Env, []LisrpValue) (LisrpValue, *LisrpError)
}

func (f *PrimitiveFunction) String() string {
	return fmt.Sprintf("<function %s>", f.Name.Id)
}

func (f *PrimitiveFunction) Type() LisrpType {
	return FunctionT
}
func (f *PrimitiveFunction) Call(env *Env, args []LisrpValue) (LisrpValue, *LisrpError) {
	return f.Code(env, args)
}
