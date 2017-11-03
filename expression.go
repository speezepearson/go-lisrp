package lisrp

type Expression interface {
	Eval(*Env) (LisrpValue, *LisrpError)
}
