package lisrp

type Expression interface {
	Eval(*Env) (Value, *LisrpError)
}
