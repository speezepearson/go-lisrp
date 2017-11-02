package lisrp

import (
	"fmt"
	"strconv"
)

type ValueType string

const (
	VTSymbol            ValueType = "Symbol"
	VTInt                         = "Int"
	VTFunction                    = "Function"
	VTPrimitiveFunction           = "PrimitiveFunction"
)

type LisrpError struct {
	Message string
}

type Env struct {
	Bindings map[string]*Value
	Parent   *Env
}

type PrimitiveFunction struct {
	Name string
	Call func([]*Value, *Env) (*Value, *LisrpError)
}

func (f *PrimitiveFunction) String() string {
	return fmt.Sprintf("<function %s>", f.Name)
}

type Function struct {
	Closure  *Env
	Name     string
	ArgNames []string
	Body     Expr
}

func (f *Function) String() string {
	return fmt.Sprintf("<function %s>", f.Name)
}

type Value struct {
	ValueType
	Symbol string
	Int    int
	Function
	PrimitiveFunction
}

func (v *Value) String() string {
	switch v.ValueType {
	case VTSymbol:
		return "'" + v.Symbol
	case VTInt:
		return strconv.Itoa(v.Int)
	case VTPrimitiveFunction:
		return v.PrimitiveFunction.String()
	case VTFunction:
		return v.Function.String()
	}
	panic(fmt.Sprintf("unknown value type %v", v.ValueType))
}

func FindBinding(name string, env *Env) (*Value, bool) {
	found := false
	for (!found) && (env != nil) {
		value, found := env.Bindings[name]
		if found {
			return value, true
		}
		env = env.Parent
	}
	return nil, false
}

func CallFunction(f Function, args []*Value, env *Env) (*Value, *LisrpError) {
	if len(f.ArgNames) != len(args) {
		return nil, &LisrpError{fmt.Sprintf("function %s expected %d args, got %d", f.Name, len(f.ArgNames), len(args))}
	}
	new_env := Env{Parent: env}
	for i, _ := range args {
		new_env.Bindings[f.ArgNames[i]] = args[i]
	}
	return Eval(f.Body, &new_env)
}

func Eval(e Expr, env *Env) (*Value, *LisrpError) {
	switch e.ExprType {
	case ETInt:
		return &Value{ValueType: VTInt, Int: e.Int}, nil
	case ETSymbol:
		value, found := FindBinding(e.Symbol, env)
		if found {
			return value, nil
		}
		return nil, &LisrpError{fmt.Sprintf("reference to unbound variable %v", e.Symbol)}
	case ETSExpr:
		if len(e.SubExprs) == 0 {
			return nil, &LisrpError{"evaluating empty s-expr"}
		}

		head, lerr := Eval(e.SubExprs[0], env)
		if lerr != nil {
			return nil, lerr
		}

		if head.ValueType != VTFunction && head.ValueType != VTPrimitiveFunction {
			return nil, &LisrpError{fmt.Sprintf("attempting to call non-function value %v", head)}
		}
		args := make([]*Value, len(e.SubExprs)-1)
		for i, _ := range args {
			args[i], lerr = Eval(e.SubExprs[i+1], env)
			if lerr != nil {
				return nil, lerr
			}
		}

		if head.ValueType == VTFunction {
			return CallFunction(head.Function, args, env)
		} else {
			return head.PrimitiveFunction.Call(args, env)
		}
	}
	panic(fmt.Sprintf("unrecognized expr type %v", e.ExprType))
}
