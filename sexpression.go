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

func (e *SExpression) Eval(env *Env) (Value, *LisrpError) {
	if len(e.SubExpressions) == 0 {
		return nil, &LisrpError{"evaluating empty s-expr"}
	}

	symbol_head, ok := e.SubExpressions[0].(*Symbol)
	if ok {
		might_be_special_or_macro := true
		for might_be_special_or_macro {
			special_form, ok := SpecialForms[*symbol_head]
			if ok {
				return special_form.Eval(env, e.SubExpressions[1:])
			}

			macro, ok := env.FindMacro(symbol_head)
			if ok {
				expansion_result := macro.Expand(e, env)
				new_e, ok := expansion_result.(*SExpression)
				if ok {
					e = new_e
				} else {
					return expansion_result.(Value), nil
				}
				continue
			}

			might_be_special_or_macro = false
		}
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