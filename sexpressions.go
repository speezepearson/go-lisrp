package lisrp

import (
	"fmt"
	"strings"
)

type Callable interface {
	Call(*Env, []Value) (Value, *LisrpError)
}

type SExpression struct {
	SubExpressions []Expression
}

func (e SExpression) String() string {
	words := make([]string, len(e.SubExpressions))
	for i, e := range e.SubExpressions {
		stringer, ok := e.(fmt.Stringer)
		if !ok {
			words[i] = "<???>"
		} else {
			words[i] = stringer.String()
		}
	}
	return fmt.Sprintf("(%s)", strings.Join(words, " "))
}

func (e SExpression) Eval(env *Env) (Value, *LisrpError) {
	if len(e.SubExpressions) == 0 {
		return nil, &LisrpError{"evaluating empty s-expr"}
	}

	head, lerr := e.SubExpressions[0].Eval(env)
	if lerr != nil {
		return nil, lerr
	}

	callable_head, ok := head.(Callable)
	if !ok {
		return nil, &LisrpError{fmt.Sprintf("attempting to call non-function value %v", head)}
	}
	args := make([]Value, len(e.SubExpressions)-1)
	for i, _ := range args {
		args[i], lerr = e.SubExpressions[i+1].Eval(env)
		if lerr != nil {
			return nil, lerr
		}
	}

	return callable_head.Call(env, args)
}
