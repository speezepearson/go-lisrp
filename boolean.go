package lisrp

type Boolean struct {
	LisrpValue bool
}

var LisrpTrue = Boolean{true}
var LisrpFalse = Boolean{false}

func (b *Boolean) Type() LisrpType {
	return BooleanT
}
func (b *Boolean) Eval(env *Env) (LisrpValue, *LisrpError) {
	return b, nil
}
