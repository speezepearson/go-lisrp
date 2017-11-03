package lisrp

type LisrpBoolean struct {
	Value bool
}

var LisrpTrue = LisrpBoolean{true}
var LisrpFalse = LisrpBoolean{false}

func (b *LisrpBoolean) Eval(env *Env) (Value, *LisrpError) {
	return b, nil
}
