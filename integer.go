package lisrp

import (
	"strconv"
)

type Integer struct {
	LisrpValue int
}

func (n Integer) String() string {
	return strconv.Itoa(n.LisrpValue)
}

func (n *Integer) Type() LisrpType {
	return IntegerT
}
func (n *Integer) Eval(env *Env) (LisrpValue, *LisrpError) {
	return n, nil
}
