package lisrp

import (
	"strconv"
)

type Integer struct {
	Value int
}

func (n Integer) String() string {
	return strconv.Itoa(n.Value)
}

func (n *Integer) Eval(env *Env) (Value, *LisrpError) {
	return n, nil
}
