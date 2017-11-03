package lisrp

type void struct{}

var VoidT = LisrpType{"void"}

func (_ *void) Type() LisrpType {
	return VoidT
}

var Void = void{}
