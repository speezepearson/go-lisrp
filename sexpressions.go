package lisrp

import (
	"fmt"
	"strings"
)

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

func (e *SExpression) MacroExpand(env *Env) interface{} {
	if len(e.SubExpressions) == 0 {
		return e
	}
	symbol, is_symbol := e.SubExpressions[0].(*Symbol)
	if !is_symbol {
		return e
	}
	macro, needs_expansion := env.FindMacro(symbol.Id)
	if !needs_expansion {
		return e
	}
	return macro.Expand(e, env)
}

func (e *SExpression) Eval(env *Env) (Value, *LisrpError) {
	if len(e.SubExpressions) == 0 {
		return nil, &LisrpError{"evaluating empty s-expr"}
	}

	var prev_val interface{}
	var new_val interface{}
	for new_val = e; prev_val != new_val; {
		new_sexpr, is_sexpression := new_val.(*SExpression)
		if !is_sexpression {
			return new_val, nil
		}
		prev_val = new_sexpr
		new_val = new_sexpr.MacroExpand(env)
	}

	// fmt.Println(e.SubExpressions[0])
	// fmt.Println(env)
	// fmt.Println(e.SubExpressions[0].Eval(env))
	// tmp, _ := e.SubExpressions[0].Eval(env)
	// fmt.Printf("foo %T\n", tmp)
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
