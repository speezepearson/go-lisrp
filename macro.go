package lisrp

type Macro interface {
	Expand(*SExpression, *Env) interface{}
}

type DefinedMacro struct {
}

type PrimitiveMacro struct {
	Name       *Symbol
	ExpandFunc func(*SExpression, *Env) interface{}
}

func (m *PrimitiveMacro) Expand(e *SExpression, env *Env) interface{} {
	return m.ExpandFunc(e, env)
}
