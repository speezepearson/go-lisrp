package lisrp

import (
	"fmt"
)

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
	head := e.SubExpressions[0].(*Symbol)
	if *head != *m.Name {
		panic(fmt.Sprintf("trying to use macro %v to expand s-expr with head %v", m.Name, head))
	}
	return m.ExpandFunc(e, env)
}
