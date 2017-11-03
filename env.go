package lisrp

type Env struct {
	Bindings      map[Symbol]Value
	MacroBindings map[Symbol]Macro
	Parent        *Env
}

func (env *Env) FindMacro(name *Symbol) (Macro, bool) {
	found := false
	for (!found) && (env != nil) {
		value, found := env.MacroBindings[*name]
		if found {
			return value, true
		}
		env = env.Parent
	}
	return nil, false
}

func (env *Env) FindBinding(name *Symbol) (Value, bool) {
	found := false
	for (!found) && (env != nil) {
		value, found := env.Bindings[*name]
		if found {
			return value, true
		}
		env = env.Parent
	}
	return nil, false
}
