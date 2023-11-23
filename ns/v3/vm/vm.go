// Package vm is responsible for the interop between legacy NoxScript VM and new NS runtime.
package vm

// Global returns a reference to a global variable in legacy NoxScript.
//
// Variables in NoxScript do not have a name, thus they are referred to by an index.
//
// For example, gvar0 in the editor corresponds to index 0 and can be accessed with vm.Global(0).
func Global(gvar int) Var {
	return Var{1, gvar}
}

// FuncVar returns a reference to a local function variable in legacy NoxScript.
//
// It's a shortcut for vm.FuncPtr(fnc).Local(vari). See Func.Local for details.
func FuncVar(fnc string, vari int) Var {
	return FuncPtr(fnc).Local(vari)
}

// Var is a reference to a legacy NoxScript variable.
type Var struct {
	fnc int
	i   int
}

// Get a variable value from legacy NoxScript.
func (vr Var) Get() Value {
	if impl == nil || vr.fnc < 0 {
		return Value{}
	}
	return Value{impl.GetFuncVar(vr.fnc, vr.i)}
}

// Set a variable value in legacy NoxScript.
func (vr Var) Set(val any) {
	if impl == nil || vr.fnc < 0 {
		return
	}
	v := toValue(val)
	impl.SetFuncVar(vr.fnc, vr.i, v.data)
}

// FuncPtr returns a function reference to a legacy NoxScript function.
func FuncPtr(fnc string) *Func {
	if impl == nil {
		return nil
	}
	fi := impl.GetFuncInd(fnc)
	if fi < 0 {
		return nil
	}
	return &Func{fnc: fi}
}

// Func is a reference to a legacy NoxScript function.
type Func struct {
	fnc int
}

// Local returns a reference to a local function variable in legacy NoxScript.
//
// Local function variables are static in NoxScript, they are not reset after the call.
// Because of this, they are frequently used to keep state between the calls.
//
// Accessing these variables is similar to Global: var0 in function Foo can be accessed with vm.FuncPtr("Foo").Local(0).
func (f *Func) Local(vari int) Var {
	if f == nil {
		return Var{-1, vari}
	}
	return Var{f.fnc, vari}
}

// Call a specified function from legacy NoxScript.
func (f *Func) Call(args ...any) Value {
	if f == nil || f.fnc < 0 || impl == nil {
		return Value{}
	}
	in := make([]uint32, len(args))
	for i, val := range args {
		in[i] = toValue(val).data
	}
	return Value{impl.CallFunc(f.fnc, in)}
}
