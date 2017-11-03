package lisrp

type Env struct {
	Bindings      map[string]Value
	MacroBindings map[string]Macro
	Parent        *Env
}

func (env *Env) FindMacro(name string) (Macro, bool) {
	found := false
	for (!found) && (env != nil) {
		value, found := env.MacroBindings[name]
		if found {
			return value, true
		}
		env = env.Parent
	}
	return nil, false
}

func (env *Env) FindBinding(name string) (Value, bool) {
	found := false
	for (!found) && (env != nil) {
		value, found := env.Bindings[name]
		if found {
			return value, true
		}
		env = env.Parent
	}
	return nil, false
}
