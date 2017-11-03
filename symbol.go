package lisrp

import (
	"fmt"
)

type Symbol struct {
	Id string
}

func (sym Symbol) String() string {
	return sym.Id
}

func (sym *Symbol) Eval(env *Env) (Value, *LisrpError) {
	value, found := env.FindBinding(sym)
	if found {
		return value, nil
	}
	return nil, &LisrpError{fmt.Sprintf("reference to unbound variable %v", sym.Id)}
}
