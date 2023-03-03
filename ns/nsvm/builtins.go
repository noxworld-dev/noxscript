package noxscript

import (
	"fmt"
	"image"
)

// TODO: should also accept a reference to Nox engine

type BuiltinFunc func(r *Runtime) (int, error)

type builtinsData struct {
	byOp     map[int]Builtin
	NamePref string
	PosOff   image.Point
}

type Builtin struct {
	Fnc           BuiltinFunc
	NeedsNamePref bool
	NeedsPosOff   bool
	ReadMore      bool
	ReadEvenMore  bool
}

func (r *Runtime) RegisterBuiltin(code int, b *Builtin) {
	if r.builtins.byOp == nil {
		r.builtins.byOp = make(map[int]Builtin)
	}
	if b == nil {
		delete(r.builtins.byOp, code)
	} else {
		if b.Fnc == nil {
			panic("function must be set")
		}
		r.builtins.byOp[code] = *b
	}
}

func (r *Runtime) resetBuiltinData() {
	r.builtins.NamePref = ""
	r.builtins.PosOff = image.Point{}
}

func (r *Runtime) BuiltinNamePref() string {
	return r.builtins.NamePref
}

func (r *Runtime) BuiltinPosOff() image.Point {
	return r.builtins.PosOff
}

func (f *Func) CallBuiltin(code int) (int, error) {
	r := f.r
	b, ok := r.builtins.byOp[code]
	if !ok {
		if r.checkPanicCompiler(code) {
			return 0, nil
		}
		err := fmt.Errorf("invalid builtin index: %d (%x)", code, code)
		return 0, &Error{Func: f.Name(), Err: err}
	}
	if f.def.NamePref == "" {
		return b.Fnc(r)
	}
	if b.NeedsNamePref {
		r.builtins.NamePref = f.def.NamePref
	}
	if b.NeedsPosOff {
		r.builtins.PosOff = f.def.PosOff
	}
	res, err := b.Fnc(r)
	r.resetBuiltinData()
	return res, err
}
