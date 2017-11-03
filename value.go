package lisrp

type LisrpType struct {
	Name string
}

var BooleanT = LisrpType{"boolean"}
var IntegerT = LisrpType{"integer"}
var SExpressionT = LisrpType{"s-expression"}
var SymbolT = LisrpType{"symbol"}

type LisrpValue interface {
	Type() LisrpType
}
