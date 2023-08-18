package noxast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"math"
	"os"
	"reflect"
	"strings"

	"github.com/noxworld-dev/noxscript/ns/asm"

	ns3 "github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	reflBool   = reflect.TypeOf(false)
	reflInt    = reflect.TypeOf(int(0))
	reflFloat  = reflect.TypeOf(float32(0))
	reflString = reflect.TypeOf("")
	reflAny    = reflect.TypeOf((*any)(nil)).Elem()
	reflObj    = reflect.TypeOf((*ns3.ObjectID)(nil)).Elem()
	reflWp     = reflect.TypeOf((*ns3.WaypointID)(nil)).Elem()
)

func importPathFor(v any) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.PkgPath()
}

type Config struct {
	DoNotOptimize bool
	DoNotFold     bool
}

func Translate(s *asm.Script, c *Config) *ast.File {
	if c == nil {
		c = &Config{}
	}
	t := &translator{
		s: s,
		c: c,
		f: &ast.File{
			Name: ast.NewIdent("script"),
		},
		strings: s.Strings,
	}
	pkg := ast.NewIdent("ns")
	t.types.nil = ast.NewIdent("nil")
	t.types.int = ast.NewIdent("int")
	t.types.float = ast.NewIdent("float32")
	t.types.string = ast.NewIdent("string")
	t.types.bool = ast.NewIdent("bool")
	t.types.true = t.builtinVar("true", reflect.TypeOf(true))
	t.types.false = t.builtinVar("false", reflect.TypeOf(false))
	t.types.Obj = &ast.SelectorExpr{Sel: ast.NewIdent("ObjectID"), X: pkg}
	t.types.Waypoint = &ast.SelectorExpr{Sel: ast.NewIdent("WaypointID"), X: pkg}
	t.imports.ns3 = pkg
	t.imports.strconv = ast.NewIdent("strconv")
	t.imports.decl = &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{},
	}
	t.f.Decls = append(t.f.Decls, t.imports.decl)
	for _, d := range builtins {
		if d == nil {
			t.builtins = append(t.builtins, nil)
			continue
		}
		id := ast.NewIdent(d.Name)
		id.Obj = &ast.Object{Name: id.Name, Data: d, Kind: ast.Fun, Type: d.Type}
		t.builtins = append(t.builtins, &ast.SelectorExpr{Sel: id, X: pkg})
	}
	t.Translate()
	return t.f
}

type TypeSet struct {
	Def   asm.VarDef
	Hints []reflect.Type
}

func (s *TypeSet) Has(t reflect.Type) bool {
	for _, h := range s.Hints {
		if h == t {
			return true
		}
	}
	return false
}

func (s *TypeSet) AllHasKind(k reflect.Kind) bool {
	for _, h := range s.Hints {
		if h.Kind() != k {
			return false
		}
	}
	return len(s.Hints) != 0
}

func (s *TypeSet) AllHasKindExcept(k reflect.Kind, exc int) bool {
	for i, h := range s.Hints {
		if i == exc {
			continue
		}
		if h.Kind() != k {
			return false
		}
	}
	return len(s.Hints) > 1
}

func (s *TypeSet) AllImplements(t reflect.Type) bool {
	if t.Kind() != reflect.Interface {
		return false
	}
	for _, h := range s.Hints {
		if !h.Implements(t) {
			return false
		}
	}
	return len(s.Hints) != 0
}

func (s *TypeSet) HasKind(k reflect.Kind) bool {
	for _, h := range s.Hints {
		if h.Kind() == k {
			return true
		}
	}
	return false
}

func (s *TypeSet) GetWithKind(k reflect.Kind) (reflect.Type, bool) {
	for _, h := range s.Hints {
		if h.Kind() == k {
			return h, true
		}
	}
	return nil, false
}

func (s *TypeSet) optimize() {
	for i := 0; i < len(s.Hints) && len(s.Hints) > 1; i++ {
		t1 := s.Hints[i]
		if s.AllImplements(t1) || (t1 == reflBool && s.AllHasKindExcept(reflect.Interface, i)) {
			s.Hints = append(s.Hints[:i], s.Hints[i+1:]...)
			i--
		}
	}
}

func (s *TypeSet) Add(types ...reflect.Type) {
	for _, t := range types {
		if t == reflAny || s.Has(t) {
			continue
		}
		if !s.AllImplements(t) {
			s.Hints = append(s.Hints, t)
		}
	}
	s.optimize()
}

func (s *TypeSet) Best() (reflect.Type, bool) {
	if len(s.Hints) == 0 {
		return nil, false
	} else if len(s.Hints) == 1 {
		return s.Hints[0], true
	} else if len(s.Hints) > 2 {
		if s.AllHasKind(reflect.Interface) && s.AllImplements(reflObj) {
			return reflObj, true
		}
		return nil, false
	}
	if s.HasKind(reflect.Bool) {
		if t, ok := s.GetWithKind(reflect.Interface); ok {
			return t, true
		}
	}
	if s.Has(reflObj) && s.Has(reflWp) {
		return reflWp, true
	}
	if s.AllHasKind(reflect.Interface) && s.AllImplements(reflObj) {
		return reflObj, true
	}
	return nil, false
}

func typeHint(x ast.Expr, types ...reflect.Type) {
	if set, ok := getType(x).(*TypeSet); ok {
		set.Add(types...)
	}
}

func typeHintExchange(x, y ast.Expr) {
	// infer down
	if types := typeHintFrom(y); len(types) != 0 {
		typeHint(x, types...)
	}
	// infer up
	if types := typeHintFrom(x); len(types) != 0 {
		typeHint(y, types...)
	}
}

func typeOf(x ast.Expr) (reflect.Type, bool) {
	types := typeHintFrom(x)
	if len(types) == 1 {
		return types[0], true
	}
	return nil, false
}

func typeHintFrom(x ast.Expr) []reflect.Type {
	switch typ := getType(x).(type) {
	case reflect.Type:
		return []reflect.Type{typ}
	case *TypeSet:
		return typ.Hints
	}
	switch x := x.(type) {
	case *ast.BasicLit:
		switch x.Kind {
		case token.STRING:
			return []reflect.Type{reflect.TypeOf("")}
		}
	case *ast.CallExpr:
		if fnc := getObj(x.Fun); fnc != nil {
			if typ, ok := fnc.Type.(reflect.Type); ok && typ.Kind() == reflect.Func {
				if typ.NumOut() == 1 {
					return []reflect.Type{typ.Out(0)}
				}
			}
			log.Printf("TODO: no hints from %q", fnc.Name)
		}
	}
	return nil
}

type translator struct {
	s     *asm.Script
	c     *Config
	f     *ast.File
	types struct {
		nil      *ast.Ident
		int      *ast.Ident
		bool     *ast.Ident
		true     *ast.Ident
		false    *ast.Ident
		float    *ast.Ident
		string   *ast.Ident
		Obj      ast.Expr
		Waypoint ast.Expr
	}
	builtins []ast.Expr
	globals  []ast.Expr
	funcs    []*ast.Ident
	strings  []string
	imports  struct {
		decl    *ast.GenDecl
		strconv *ast.Ident
		ns3     *ast.Ident
	}
}

var nameReplacer = strings.NewReplacer(
	" ", "_",
	"-", "_",
)

func cleanGoName(name string) string {
	name = strings.TrimSpace(name)
	if i := strings.IndexByte(name, ':'); i > 0 {
		name = name[:i]
	}
	name = nameReplacer.Replace(name)
	if name[0] >= '0' && name[0] <= '9' {
		name = "_" + name
	}
	return name
}

func (t *translator) Translate() {
	for i, fnc := range t.s.Funcs {
		switch i {
		case 0, 1:
			t.funcs = append(t.funcs, ast.NewIdent("init"))
		default:
			name := cleanGoName(fnc.Name)
			t.funcs = append(t.funcs, ast.NewIdent(name))
		}
	}
	for i, fnc := range t.s.Funcs {
		id := t.funcs[i]
		switch i {
		case 0:
			t.translateGlobal0(id, fnc)
		case 1:
			t.translateGlobal1(id, fnc)
		default:
			d := &ast.FuncDecl{Name: id}
			t.translateFunc(d, fnc, false)
			t.f.Decls = append(t.f.Decls, d)
		}
	}
	t.inferTypes()
	t.fixBoolAndNil()
	t.addOnEvent()
	t.updateImports()
}

func (t *translator) translateGlobal0(id *ast.Ident, f asm.FuncDef) {
	if len(f.Code) > 1 {
		d := &ast.FuncDecl{Name: id}
		t.translateFunc(d, f, true)
		t.f.Decls = append(t.f.Decls, d)
	}
}

func (t *translator) builtinVarPkg(name string, typ reflect.Type) ast.Expr {
	id := ast.NewIdent(name)
	id.Obj = &ast.Object{Kind: ast.Var, Name: id.Name, Type: typ}
	return &ast.SelectorExpr{Sel: id, X: t.imports.ns3}
}

func (t *translator) builtinVar(name string, typ reflect.Type) *ast.Ident {
	id := ast.NewIdent(name)
	id.Obj = &ast.Object{Kind: ast.Var, Name: id.Name, Type: typ}
	return id
}

func (t *translator) translateGlobal1(id *ast.Ident, f asm.FuncDef) {
	t.globals = append(t.globals,
		&ast.CallExpr{Fun: t.builtins[asm.BuiltinGetTrigger]},
		&ast.CallExpr{Fun: t.builtins[asm.BuiltinGetCaller]},
		t.types.true,
		t.types.false,
	)
	if len(f.Vars) > 4 {
		d := &ast.GenDecl{Tok: token.VAR}
		for i, v := range f.Vars[4:] {
			id := ast.NewIdent(fmt.Sprintf("gvar%d", 4+i))
			sp := &ast.ValueSpec{Names: []*ast.Ident{id}, Type: t.types.int}
			if v.Size > 1 {
				sp.Type = &ast.ArrayType{Len: intLit(v.Size), Elt: sp.Type}
			}
			id.Obj = &ast.Object{Kind: ast.Var, Name: id.Name, Decl: sp, Type: &TypeSet{Def: v}}
			d.Specs = append(d.Specs, sp)
			t.globals = append(t.globals, id)
		}
		t.f.Decls = append(t.f.Decls, d)
	}
	if len(f.Code) > 1 {
		d := &ast.FuncDecl{Name: id}
		t.translateFunc(d, f, true)
		t.f.Decls = append(t.f.Decls, d)
	}
}

func (t *translator) translateFunc(d *ast.FuncDecl, f asm.FuncDef, global bool) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("panic when translating %q: %v", f.Name, r))
		}
	}()
	d.Type = &ast.FuncType{
		Params: &ast.FieldList{},
	}
	d.Body = &ast.BlockStmt{}
	if global {
		t.translateCode(d.Body, f.Return > 0, t.globals, f.Code)
		return
	}
	if f.Return > 0 {
		d.Type.Results = &ast.FieldList{List: []*ast.Field{
			{Type: t.types.int},
		}}
	}
	vars := make([]ast.Expr, 0, len(f.Vars))
	for i, v := range f.Vars[:f.Args] {
		id := ast.NewIdent(fmt.Sprintf("a%d", i+1))
		fld := &ast.Field{Names: []*ast.Ident{id}, Type: t.types.int}
		if v.Size > 1 {
			fld.Type = &ast.ArrayType{Len: intLit(v.Size), Elt: fld.Type}
		}
		id.Obj = &ast.Object{Kind: ast.Var, Name: id.Name, Decl: fld, Type: &TypeSet{Def: v}}
		d.Type.Params.List = append(d.Type.Params.List, fld)
		vars = append(vars, id)
	}
	if f.Args < len(f.Vars) {
		vd := &ast.GenDecl{Tok: token.VAR}
		for i, v := range f.Vars[f.Args:] {
			id := ast.NewIdent(fmt.Sprintf("v%d", i))
			sp := &ast.ValueSpec{Names: []*ast.Ident{id}, Type: t.types.int}
			if v.Size > 1 {
				sp.Type = &ast.ArrayType{Len: intLit(v.Size), Elt: sp.Type}
			}
			id.Obj = &ast.Object{Kind: ast.Var, Name: id.Name, Decl: sp, Type: &TypeSet{Def: v}}
			vd.Specs = append(vd.Specs, sp)
			vars = append(vars, id)
		}
		d.Body.List = append(d.Body.List, &ast.DeclStmt{Decl: vd})
	}
	t.translateCode(d.Body, f.Return > 0, vars, f.Code)
}

type temporary struct{}

func (t *translator) translateCode(d *ast.BlockStmt, ret bool, vars []ast.Expr, code []uint32) {
	list, err := asm.Decode(code)
	if err != nil {
		d.List = append(d.List, &ast.ExprStmt{X: ast.NewIdent("/* " + err.Error() + " */")})
		return
	}
	var (
		codeOff   int
		lastTmp   int
		stack     []ast.Expr
		work      []ast.Stmt
		nextLabel *ast.Ident
		labels    = make(map[int]*ast.Ident)
		debug     = false
		debugBuf  bytes.Buffer
	)
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, debugBuf.String())
			panic(r)
		}
	}()
	for _, v := range list {
		jmp, ok := v.(asm.Jump)
		if !ok {
			continue
		}
		if l := labels[int(jmp.Off)]; l == nil {
			l = ast.NewIdent(fmt.Sprintf("LABEL%d", len(labels)+1))
			labels[int(jmp.Off)] = l
		}
	}
	pop := func() ast.Expr {
		n := len(stack)
		x := stack[n-1]
		if debug {
			fmt.Fprintf(&debugBuf, "\t\tPOP %s\n", printExpr(x))
		}
		stack = stack[:n-1]
		return x
	}
	maybePop := func() ast.Expr {
		if len(stack) == 0 {
			if debug {
				fmt.Fprintf(&debugBuf, "\t\tPOP 0\n")
			}
			return intLit(0)
		}
		return pop()
	}
	push := func(x ast.Expr) {
		if debug {
			fmt.Fprintf(&debugBuf, "\t\tPUSH %s\n", printExpr(x))
		}
		stack = append(stack, x)
	}
	stmt := func(s ast.Stmt) {
		if nextLabel != nil {
			s = &ast.LabeledStmt{Label: nextLabel, Stmt: s}
			nextLabel = nil
		}
		work = append(work, s)
	}
	tmpVar := func(x ast.Expr) *ast.Ident {
		id := ast.NewIdent(fmt.Sprintf("r%d", lastTmp))
		id.Obj = &ast.Object{Name: id.Name, Kind: ast.Var, Data: temporary{}}
		if types := typeHintFrom(x); len(types) == 1 {
			id.Obj.Type = types[0]
		} else if len(types) != 0 {
			id.Obj.Type = &TypeSet{Hints: types}
		}
		lastTmp++
		stmt(&ast.AssignStmt{Lhs: []ast.Expr{id}, Tok: token.DEFINE, Rhs: []ast.Expr{x}})
		return id
	}
	for _, v := range list {
		if debug {
			fmt.Fprintf(&debugBuf, "%d:  %s\n", codeOff, v)
		}
		switch v := v.(type) {
		default:
			panic(v.OpCode().String())
		case asm.Return:
			switch v.Op {
			default:
				panic(v.Op.String())
			case asm.OpReturn, asm.OpReturn0:
				if ret {
					rv := maybePop()
					if rn, ok := rv.(*ast.Ident); ok {
						switch rn.Name {
						case "false":
							rv = intLit(0)
						case "true":
							rv = intLit(1)
						}
					}
					stmt(&ast.ReturnStmt{Results: []ast.Expr{rv}})
				} else {
					stmt(&ast.ReturnStmt{})
				}
			}
		case asm.Jump:
			s := &ast.BranchStmt{Tok: token.GOTO, Label: labels[int(v.Off)]}
			switch v.Op {
			default:
				panic(v.Op.String())
			case asm.OpJump:
				stmt(s)
			case asm.OpJumpIf:
				stmt(&ast.IfStmt{Cond: condition(pop()), Body: &ast.BlockStmt{List: []ast.Stmt{s}}})
			case asm.OpJumpIfNot:
				stmt(&ast.IfStmt{Cond: not(condition(pop())), Body: &ast.BlockStmt{List: []ast.Stmt{s}}})
			}
		case asm.CallScript:
			id := t.funcs[v.Index]
			d := t.s.Funcs[v.Index]
			x := &ast.CallExpr{Fun: id}
			for range d.Vars[:d.Args] {
				a := pop()
				x.Args = append(x.Args, a)
			}
			if d.Return > 0 {
				push(tmpVar(x))
			} else {
				stmt(&ast.ExprStmt{X: x})
			}
		case asm.CallBuiltin:
			rt, call := t.callBuiltin(v.Index)
			args := make([]ast.Expr, rt.NumIn())
			for i := rt.NumIn() - 1; i >= 0; i-- {
				a := pop()
				at := rt.In(i)
				typeHint(a, at)
				if val, ok := asInt(a); ok && val == 0 && at.Kind() == reflect.Interface {
					a = t.types.nil
				}
				args[i] = a
			}
			x := call(args)
			if rt.NumOut() > 0 {
				push(tmpVar(x))
			} else {
				stmt(&ast.ExprStmt{X: x})
			}
		case asm.Push:
			switch v.Op {
			default:
				panic(v.Op.String())
			case asm.OpPushInt:
				push(intLit(int(v.Val)))
			case asm.OpPushFloat:
				push(floatLit(math.Float32frombits(uint32(v.Val))))
			case asm.OpPushString:
				push(stringLit(t.strings[v.Val]))
			}
		case asm.LoadVar:
			switch v.Op {
			default:
				panic(v.Op.String())
			case asm.OpLoadVarInt, asm.OpLoadVarFloat, asm.OpLoadVarString, asm.OpLoadVarPtr:
				var x ast.Expr
				if v.IsGlobal != 0 {
					if v.Index < 0 || int(v.Index) >= len(t.globals) {
						x = ast.NewIdent(fmt.Sprintf("gvarUnk%d", v.Index))
					} else {
						x = t.globals[v.Index]
					}
				} else {
					if v.Index < 0 || int(v.Index) >= len(vars) {
						x = ast.NewIdent(fmt.Sprintf("unk%d", v.Index))
					} else {
						x = vars[v.Index]
					}
				}
				switch v.Op {
				case asm.OpLoadVarFloat:
					typeHint(x, reflFloat)
				case asm.OpLoadVarString:
					typeHint(x, reflString)
				}
				if v.Op == asm.OpLoadVarPtr {
					if sz, ok := arrayLen(x); ok {
						push(intLit(sz))
					} else if debug {
						fmt.Fprintf(&debugBuf, "(not pushing size: type %T)", getType(x))
					}
					push(intLit(int(v.IsGlobal)))
					push(takeAddr(x))
				} else {
					push(x)
				}
			}
		case asm.Index:
			switch v.Op {
			default:
				panic(v.Op.String())
			case asm.OpIndexInt, asm.OpIndexFloat, asm.OpIndexPtr:
				off := pop()
				exp := pop()
				switch v.Op {
				case asm.OpIndexFloat:
					typeHint(exp, reflFloat)
				}
				isGlobalX := pop()
				szX := pop() // TODO: check

				isGlobal, gok := asInt(isGlobalX)
				ind, isInd := asInt(exp)
				indU, isPtr := exp.(*ast.UnaryExpr)
				isPtr = isPtr && indU.Op == token.AND

				var rhs ast.Expr
				if gok && (isInd || isPtr) {
					if !isInd {
						rhs = indU.X
					} else if isGlobal != 0 {
						if ind < 0 || ind >= len(t.globals) {
							rhs = ast.NewIdent(fmt.Sprintf("gvarUnk%d", ind))
						} else {
							rhs = t.globals[ind]
						}
					} else {
						if ind < 0 || ind >= len(vars) {
							rhs = ast.NewIdent(fmt.Sprintf("unk%d", ind))
						} else {
							rhs = vars[ind]
						}
					}
					rhs = &ast.IndexExpr{X: rhs, Index: off}
				} else {
					debug = true
					rhs = &ast.CallExpr{Fun: ast.NewIdent("__dynamic_var_get"), Args: []ast.Expr{isGlobalX, exp, off, szX}}
				}
				if v.Op == asm.OpIndexPtr {
					push(isGlobalX)
					push(takeAddr(rhs))
				} else {
					push(rhs)
				}
			}
		case asm.StoreVar:
			rhs := pop()
			ptrX := pop()
			isGlobalX := pop()

			isGlobal, gok := asInt(isGlobalX)
			ind, isInd := asInt(ptrX)
			indU, isPtr := ptrX.(*ast.UnaryExpr)
			isPtr = isPtr && indU.Op == token.AND

			var op = token.ASSIGN
			switch v.Op {
			default:
				panic(v.Op.String())
			case asm.OpStoreInt, asm.OpStoreFloat, asm.OpStoreString:
				// assign
			case asm.OpStoreIntAdd, asm.OpStoreFloatAdd, asm.OpStoreStringAdd:
				op = token.ADD_ASSIGN
			case asm.OpStoreIntSub, asm.OpStoreFloatSub:
				op = token.SUB_ASSIGN
			case asm.OpStoreIntMul, asm.OpStoreFloatMul:
				op = token.MUL_ASSIGN
			case asm.OpStoreIntDiv, asm.OpStoreFloatDiv:
				op = token.QUO_ASSIGN
			case asm.OpStoreIntMod:
				op = token.REM_ASSIGN
			}
			push(rhs)
			if gok && (isInd || isPtr) {
				var lhs ast.Expr
				if isPtr {
					lhs = indU.X
				} else if isGlobal != 0 {
					if ind < 0 || ind >= len(t.globals) {
						lhs = ast.NewIdent(fmt.Sprintf("gvarUnk%d", ind))
					} else {
						lhs = t.globals[ind]
					}
				} else {
					if ind < 0 || ind >= len(vars) {
						lhs = ast.NewIdent(fmt.Sprintf("unk%d", ind))
					} else {
						lhs = vars[ind]
					}
				}
				switch v.Op {
				case asm.OpStoreFloat, asm.OpStoreFloatAdd, asm.OpStoreFloatSub, asm.OpStoreFloatMul, asm.OpStoreFloatDiv:
					typeHint(lhs, reflFloat)
					typeHint(rhs, reflFloat)
				case asm.OpStoreString:
					typeHint(lhs, reflString)
					typeHint(rhs, reflString)
				case asm.OpStoreIntAdd, asm.OpStoreIntSub, asm.OpStoreIntMul, asm.OpStoreIntDiv, asm.OpStoreIntMod:
					typeHint(lhs, reflInt)
					typeHint(rhs, reflInt)
				}
				typeHintExchange(lhs, rhs)
				stmt(&ast.AssignStmt{Lhs: []ast.Expr{lhs}, Tok: op, Rhs: []ast.Expr{rhs}})
			} else {
				debug = true
				stmt(&ast.ExprStmt{X: &ast.CallExpr{Fun: ast.NewIdent("__dynamic_var_set"), Args: []ast.Expr{isGlobalX, ptrX, rhs}}})
			}
		case asm.UnaryOp:
			x := pop()
			var op token.Token
			switch v.Op {
			default:
				panic(v.Op.String())
			case asm.OpBoolNot:
				typeHint(x, reflBool)
				op = token.NOT
			case asm.OpIntNot:
				op = token.XOR
			case asm.OpIntNeg:
				typeHint(x, reflInt)
				op = token.SUB
			case asm.OpFloatNeg:
				typeHint(x, reflFloat)
				op = token.SUB
			}
			push(&ast.UnaryExpr{X: x, Op: op})
		case asm.BinaryOp:
			rhs := pop()
			lhs := pop()
			switch v.Op {
			case asm.OpIntEq, asm.OpIntNeq:
				typeHintExchange(lhs, rhs)
			case asm.OpBoolAnd, asm.OpBoolOr:
				typeHint(rhs, reflBool)
				typeHint(lhs, reflBool)
			case asm.OpStringAdd, asm.OpStringEq, asm.OpStringNeq, asm.OpStringLt, asm.OpStringGt, asm.OpStringLte, asm.OpStringGte:
				typeHint(rhs, reflString)
				typeHint(lhs, reflString)
			case asm.OpFloatAdd, asm.OpFloatSub, asm.OpFloatMul, asm.OpFloatDiv,
				asm.OpFloatEq, asm.OpFloatNeq, asm.OpFloatLt, asm.OpFloatGt, asm.OpFloatLte, asm.OpFloatGte:
				typeHint(rhs, reflFloat)
				typeHint(lhs, reflFloat)
			case asm.OpIntAdd, asm.OpIntSub, asm.OpIntMul, asm.OpIntDiv, asm.OpIntMod,
				asm.OpIntLt, asm.OpIntGt, asm.OpIntLte, asm.OpIntGte,
				asm.OpIntXOr, asm.OpIntLSh, asm.OpIntRSh:
				typeHint(rhs, reflInt)
				typeHint(lhs, reflInt)
			}
			var op token.Token
			switch v.Op {
			default:
				panic(v.Op.String())
			case asm.OpIntAdd, asm.OpFloatAdd, asm.OpStringAdd:
				op = token.ADD
			case asm.OpIntSub, asm.OpFloatSub:
				op = token.SUB
			case asm.OpIntMul, asm.OpFloatMul:
				op = token.MUL
			case asm.OpIntDiv, asm.OpFloatDiv:
				op = token.QUO
			case asm.OpIntMod:
				op = token.REM
			case asm.OpIntAnd:
				op = token.AND
			case asm.OpIntOr:
				op = token.OR
				swap := 0
				if lhs, ok := lhs.(*ast.BinaryExpr); ok && isLogicalOp(lhs.Op) {
					swap++
				}
				if rhs, ok := rhs.(*ast.BinaryExpr); ok && isLogicalOp(rhs.Op) {
					swap++
				}
				if swap == 2 {
					op = token.LOR
				}
			case asm.OpIntXOr:
				op = token.XOR
			case asm.OpIntLSh:
				op = token.SHL
			case asm.OpIntRSh:
				op = token.SHR
			case asm.OpIntEq, asm.OpFloatEq, asm.OpStringEq:
				op = token.EQL
			case asm.OpIntNeq, asm.OpFloatNeq, asm.OpStringNeq:
				op = token.NEQ
			case asm.OpIntLt, asm.OpFloatLt, asm.OpStringLt:
				op = token.LSS
			case asm.OpIntGt, asm.OpFloatGt, asm.OpStringGt:
				op = token.GTR
			case asm.OpIntLte, asm.OpFloatLte, asm.OpStringLte:
				op = token.LEQ
			case asm.OpIntGte, asm.OpFloatGte, asm.OpStringGte:
				op = token.GEQ
			case asm.OpBoolAnd:
				op = token.LAND
			case asm.OpBoolOr:
				op = token.LOR
			}
			push(&ast.BinaryExpr{X: lhs, Op: op, Y: rhs})
		}
		codeOff += v.Len()
		if l := labels[codeOff]; l != nil {
			nextLabel = l
		}
	}
	d.List = append(d.List, work...)
	t.simplifyCode(d, ret)
	if debug {
		var buf bytes.Buffer
		buf.WriteString("/*\n")
		buf.Write(debugBuf.Bytes())
		buf.WriteString("*/")
		d.List = append(d.List, &ast.ExprStmt{X: ast.NewIdent(buf.String())})
	}
}

func isLogicalOp(op token.Token) bool {
	switch op {
	case token.EQL, token.NEQ, token.LSS, token.GTR, token.LEQ, token.GEQ, token.LAND, token.LOR:
		return true
	}
	return false
}

func (t *translator) simplifyCode(d *ast.BlockStmt, ret bool) {
	if !t.c.DoNotOptimize {
		t.removeSingleDefines(d)
	}
	if !t.c.DoNotFold {
		t.makeSwitches(d)
		t.makeLoops(d)
		if !ret {
			t.removeLastReturn(d)
		}
	}
	if !t.c.DoNotOptimize {
		t.removeUnusedDefines(d)
		t.removeUnusedLabels(d)
		t.fixUnusedVars(d)
	}
}

func (t *translator) removeSingleDefines(d *ast.BlockStmt) {
	for i := 0; i < len(d.List)-1; i++ {
		if i < 0 {
			continue
		}
		var label *ast.Ident
		df, ok := d.List[i].(*ast.AssignStmt)
		if !ok {
			if lbl, lok := d.List[i].(*ast.LabeledStmt); lok {
				label = lbl.Label
				df, ok = lbl.Stmt.(*ast.AssignStmt)
			}
		}
		if !ok || df.Tok != token.DEFINE {
			continue
		}
		id, ok := df.Lhs[0].(*ast.Ident)
		if !ok {
			continue
		}
		_, ok = id.Obj.Data.(temporary)
		if !ok {
			continue
		}
		used := false
		for j := i + 2; j < len(d.List); j++ {
			used = used || cntUsages(d.List[j], id) > 0
		}
		if used {
			continue
		}
		val := df.Rhs[0]
		usages := cntUsages(d.List[i+1], id)
		if usages > 1 {
			continue
		}
		effectSafe := false
		if d2, ok := d.List[i+1].(*ast.AssignStmt); ok && len(d2.Rhs) == 1 && d2.Rhs[0] == id {
			effectSafe = true
		}
		if !effectSafe && hasSideEffects(val) {
			continue
		}
		replaced := replace(d.List[i+1], id, val)
		if usages != replaced {
			continue
		}
		d.List = append(d.List[:i], d.List[i+1:]...)
		if label != nil {
			d.List[i] = &ast.LabeledStmt{Label: label, Stmt: d.List[i]}
		}
		i -= 2 // check previous again
	}
}
func (t *translator) makeSwitches(d *ast.BlockStmt) {
	if true {
		return
	}
	for i := 0; i < len(d.List); i++ {
		// It usually starts with a temporary var assignment, which will be used in switch.
		var label *ast.Ident
		init, ok := d.List[i].(*ast.AssignStmt)
		if !ok {
			l, lok := d.List[i].(*ast.LabeledStmt)
			if !lok {
				continue
			}
			label = l.Label
			init, ok = l.Stmt.(*ast.AssignStmt)
		}
		if !ok || init.Tok != token.ASSIGN {
			continue
		}

		vr := init.Lhs[0]
		// Then N if statements should follow. Each has a single goto statement.
		var (
			cases   []ast.Expr
			targets []*ast.Ident
		)
		for _, st := range d.List[i+1:] {
			ifs, ok := st.(*ast.IfStmt)
			if !ok || ifs.Init != nil || ifs.Else != nil || len(ifs.Body.List) != 1 {
				break
			}
			cond, ok := ifs.Cond.(*ast.BinaryExpr)
			if !ok || cond.Op != token.EQL || cond.X != vr {
				break
			}
			targ, ok := ifs.Body.List[0].(*ast.BranchStmt)
			if !ok || targ.Tok != token.GOTO {
				break
			}
			cases = append(cases, cond.Y)
			targets = append(targets, targ.Label)
		}
		if len(cases) == 0 {
			continue
		}
		// It ends with another goto statement (a default branch of switch).
		def, ok := d.List[i+1+len(cases)].(*ast.BranchStmt)
		if !ok || def.Tok != token.GOTO {
			continue
		}
		// First we will just combine everything that we already found, and then we will try to simplify it.
		sw := &ast.SwitchStmt{
			Init: init, Tag: vr,
			Body: &ast.BlockStmt{},
		}
		for j := range cases {
			sw.Body.List = append(sw.Body.List, &ast.CaseClause{
				List: []ast.Expr{cases[j]},
				Body: []ast.Stmt{&ast.BranchStmt{Tok: token.GOTO, Label: targets[j]}},
			})
		}
		sw.Body.List = append(sw.Body.List, &ast.CaseClause{
			Body: []ast.Stmt{&ast.BranchStmt{Tok: token.GOTO, Label: def.Label}},
		})
		var st ast.Stmt = sw
		if label != nil {
			st = &ast.LabeledStmt{Label: label, Stmt: st}
		}
		d.List[i] = st
		d.List = append(d.List[:i+1], d.List[i+1+len(cases)+1:]...)
		//continue

		// Now try to find switch case bodies and move them inside.
		off := 1
	casesLoop:
		for j := 0; j < len(cases)+1 && i+off < len(d.List); j++ {
			l, ok := d.List[i+off].(*ast.LabeledStmt)
			if !ok {
				break
			}
			if cntUsages(&ast.BlockStmt{List: d.List[:i]}, l.Label)+
				cntUsages(&ast.BlockStmt{List: d.List[i+1 : i+off]}, l.Label)+
				cntUsages(&ast.BlockStmt{List: d.List[i+off+1:]}, l.Label) != 0 {
				break
			}
			off++
			var body []ast.Stmt
			body = append(body, l.Stmt)
			if st, ok := l.Stmt.(*ast.BranchStmt); ok && st.Tok == token.GOTO {
				sw.Body.List[j].(*ast.CaseClause).Body = body
				continue
			}
			for i+off < len(d.List) {
				st := d.List[i+off]
				off++
				body = append(body, st)
				if st, ok := st.(*ast.BranchStmt); ok && st.Tok == token.GOTO {
					sw.Body.List[j].(*ast.CaseClause).Body = body
					continue casesLoop
				}
			}
		}
		if off > 1 {
			d.List = append(d.List[:i+1], d.List[i+off:]...)
		}

		// If we find label right after the switch and corresponding goto in cases, we can remove gotos.
		if i+1 < len(d.List) {
			if l, ok := d.List[i+1].(*ast.LabeledStmt); ok {
				for j := range sw.Body.List {
					c := sw.Body.List[j].(*ast.CaseClause)
					n := len(c.Body)
					if gt, ok := c.Body[n-1].(*ast.BranchStmt); ok && gt.Tok == token.GOTO && gt.Label == l.Label {
						c.Body = c.Body[:n-1]
					}
				}
			}
		}

		// If default branch is now empty - remove it
		if c := sw.Body.List[len(sw.Body.List)-1].(*ast.CaseClause); len(c.List) == 0 && len(c.Body) == 0 {
			sw.Body.List = sw.Body.List[:len(sw.Body.List)-1]
		}
	}
}
func (t *translator) makeLoops(d *ast.BlockStmt) {
loops:
	for i := 0; i < len(d.List); i++ {
		l, ok := d.List[i].(*ast.LabeledStmt)
		if !ok {
			continue
		}
		for j := i + 1; j < len(d.List); j++ {
			jmp, ok := d.List[j].(*ast.BranchStmt)
			if ok && jmp.Tok == token.GOTO && jmp.Label == l.Label {
				fr := &ast.ForStmt{Body: &ast.BlockStmt{
					List: append([]ast.Stmt{}, d.List[i:j]...),
				}}
				fr.Body.List[0] = l.Stmt
				d.List[i] = &ast.LabeledStmt{Label: l.Label, Stmt: fr}
				d.List = append(d.List[:i+1], d.List[j+1:]...)
				i--
				continue loops
			}
		}
	}
}
func (t *translator) removeLastReturn(d *ast.BlockStmt) {
	if n := len(d.List); n != 0 {
		if _, ok := d.List[n-1].(*ast.ReturnStmt); ok {
			d.List = d.List[:n-1]
		}
	}
}
func (t *translator) removeUnusedDefines(d *ast.BlockStmt) {
	for i := 0; i < len(d.List); i++ {
		df, ok := d.List[i].(*ast.AssignStmt)
		if !ok || df.Tok != token.DEFINE {
			continue
		}
		id, ok := df.Lhs[0].(*ast.Ident)
		if !ok {
			continue
		}
		_, ok = id.Obj.Data.(temporary)
		if !ok {
			continue
		}
		used := false
		for j := i + 1; j < len(d.List); j++ {
			used = used || cntUsages(d.List[j], id) > 0
		}
		if !used {
			d.List[i] = &ast.ExprStmt{X: df.Rhs[0]}
		}
	}
}
func (t *translator) removeUnusedLabels(root *ast.BlockStmt) {
	ast.Inspect(root, func(n ast.Node) bool {
		if d, ok := n.(*ast.BlockStmt); ok {
			for i := 0; i < len(d.List); i++ {
				l, ok := d.List[i].(*ast.LabeledStmt)
				if !ok {
					continue
				}
				if cntUsages(root, l.Label) == 1 {
					d.List[i] = l.Stmt
				}
			}
		}
		return true
	})
}
func (t *translator) fixUnusedVars(d *ast.BlockStmt) {
	for i := 0; i < len(d.List); i++ {
		st, ok := d.List[i].(*ast.DeclStmt)
		if !ok {
			continue
		}
		decl, ok := st.Decl.(*ast.GenDecl)
		if !ok || decl.Tok != token.VAR {
			continue
		}
		for si := 0; si < len(decl.Specs); si++ {
			sp, ok := decl.Specs[si].(*ast.ValueSpec)
			if !ok {
				continue
			}
			for ni := 0; ni < len(sp.Names); ni++ {
				id := sp.Names[ni]
				usages := cntUsages(d, id)
				if usages == 1 {
					sp.Names = append(sp.Names[:ni], sp.Names[ni+1:]...)
					ni--
					continue
				}
				found := 1
				ast.Inspect(d, func(n ast.Node) bool {
					if st, ok := n.(*ast.AssignStmt); ok && st.Lhs[0] == id {
						found++
					}
					return true
				})
				if found == usages {
					d.List = append(d.List, nil)
					copy(d.List[i+2:], d.List[i+1:])
					d.List[i+1] = &ast.AssignStmt{
						Lhs: []ast.Expr{ast.NewIdent("_")},
						Tok: token.ASSIGN,
						Rhs: []ast.Expr{id},
					}
				}
			}
			if len(sp.Names) == 0 {
				decl.Specs = append(decl.Specs[:si], decl.Specs[si+1:]...)
				si--
			}
		}
		if len(decl.Specs) == 0 {
			d.List = append(d.List[:i], d.List[i+1:]...)
			i--
		}
	}
}

func (t *translator) inferTypeOf(v ast.Expr) {
	obj := getObj(v)
	if obj == nil {
		return
	}
	ts, ok := obj.Type.(*TypeSet)
	if !ok || len(ts.Hints) == 0 {
		return
	}
	//log.Printf("type hints for %q: %v", obj.Name, ts.Hints)
	rt, ok := ts.Best()
	if !ok {
		log.Printf("multiple hints for %q: %v", obj.Name, ts.Hints)
		return
	}
	ts.Hints = []reflect.Type{rt}
	obj.Type = rt
	var (
		typ  ast.Expr
		pref = "gvar"
	)
	switch rt {
	case reflBool:
		typ = t.types.bool
		pref = "flag"
	case reflInt:
		typ = t.types.int
		pref = "ivar"
	case reflString:
		typ = t.types.string
		pref = "str"
	case reflFloat:
		typ = t.types.float
		pref = "fvar"
	case reflObj:
		typ = t.types.Obj
		pref = "obj"
	case reflWp:
		typ = t.types.Waypoint
		pref = "wp"
	default:
		if rt.PkgPath() != reflect.TypeOf((*ns3.ObjectID)(nil)).Elem().PkgPath() {
			log.Printf("unsupported hint for %q: %v", obj.Name, rt)
			return
		}
		typ = &ast.SelectorExpr{Sel: ast.NewIdent(rt.Name()), X: t.imports.ns3}
	}
	name := obj.Name
	if strings.HasPrefix(name, "gvar") {
		name = pref + strings.TrimPrefix(name, "gvar")
	}
	if ts.Def.Size > 1 {
		typ = &ast.ArrayType{Len: intLit(ts.Def.Size), Elt: typ}
	}
	switch decl := obj.Decl.(type) {
	case *ast.Field:
		obj.Name = name
		decl.Names[0].Name = name
		decl.Type = typ
	case *ast.ValueSpec:
		obj.Name = name
		decl.Names[0].Name = name
		decl.Type = typ
	default:
		log.Printf("can't set inferred type for %q to %T", obj.Name, typ)
	}
}

func (t *translator) inferTypes() {
	for _, v := range t.globals {
		t.inferTypeOf(v)
	}
	ast.Inspect(t.f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.FuncDecl:
			for _, a := range n.Type.Params.List {
				t.inferTypeOf(a.Names[0])
			}
		case *ast.ValueSpec:
			for _, id := range n.Names {
				t.inferTypeOf(id)
			}
		}
		return true
	})
}
func (t *translator) fixBoolAndNil() {
	ast.Inspect(t.f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.IfStmt:
			switch x := n.Cond.(type) {
			case *ast.UnaryExpr:
				switch x.Op {
				case token.NOT:
					rt, _ := typeOf(x.X)
					if rt != nil && rt.Kind() == reflect.Interface {
						n.Cond = &ast.BinaryExpr{
							X: x.X, Op: token.EQL, Y: t.types.nil,
						}
					} else if rt == nil {
						if _, ok := x.X.(*ast.Ident); ok {
							n.Cond = &ast.BinaryExpr{
								X: x.X, Op: token.EQL, Y: intLit(0),
							}
						}
					}
				}
			default:
				rt, _ := typeOf(n.Cond)
				if rt != nil && rt.Kind() == reflect.Interface {
					n.Cond = &ast.BinaryExpr{
						X: n.Cond, Op: token.NEQ, Y: t.types.nil,
					}
				} else if rt == nil {
					if _, ok := n.Cond.(*ast.Ident); ok {
						n.Cond = &ast.BinaryExpr{
							X: n.Cond, Op: token.NEQ, Y: intLit(0),
						}
					}
				}
			}
		case *ast.AssignStmt:
			switch n.Tok {
			case token.ASSIGN, token.DEFINE:
				lhs, rhs := n.Lhs[0], n.Rhs[0]
				if val, ok := asInt(rhs); ok && val == 0 {
					rt, _ := typeOf(lhs)
					if rt != nil && rt.Kind() == reflect.Interface {
						n.Rhs[0] = t.types.nil
					}
				}
			}
		case *ast.BinaryExpr:
			switch n.Op {
			case token.EQL, token.NEQ:
				lhs, rhs := n.X, n.Y
				if val, ok := asInt(rhs); ok && val == 0 {
					rt, _ := typeOf(lhs)
					if rt != nil && rt.Kind() == reflect.Interface {
						n.Y = t.types.nil
					}
				}
			}
		}
		return true
	})
}
func (t *translator) addOnEvent() {
	events := map[string]struct{}{
		"MapInitialize": {},
		"MapEntry":      {},
		"MapExit":       {},
		"MapShutdown":   {},
		"PlayerDeath":   {},
	}
	var list []ast.Stmt
	for _, f := range t.funcs {
		_, ok := events[f.Name]
		if !ok {
			continue
		}
		// case "<event>":
		// 		<func>()
		list = append(list, &ast.CaseClause{
			List: []ast.Expr{
				stringLit(f.Name),
			},
			Body: []ast.Stmt{
				&ast.ExprStmt{X: &ast.CallExpr{Fun: f}},
			},
		})
	}
	if len(list) != 0 {
		// func OnEvent(typ string) {
		// 		switch typ {
		//		<list>
		//		}
		// }
		arg := ast.NewIdent("typ")
		t.f.Decls = append(t.f.Decls, &ast.FuncDecl{
			Name: ast.NewIdent("OnEvent"),
			Type: &ast.FuncType{Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							arg,
						},
						Type: ast.NewIdent("string"),
					},
				},
			}},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.SwitchStmt{
					Tag:  arg,
					Body: &ast.BlockStmt{List: list},
				},
			}},
		})
	}
}
func (t *translator) updateImports() {
	var used struct {
		strconv int
		ns3     int
	}
	ast.Inspect(t.f, func(n ast.Node) bool {
		switch n {
		case t.imports.strconv:
			used.strconv++
		case t.imports.ns3:
			used.ns3++
		}
		return true
	})
	var add []string
	if used.strconv != 0 {
		add = append(add, "strconv")
	}
	if used.ns3 != 0 {
		add = append(add, importPathFor((*ns3.ObjectID)(nil)))
	}
	for _, path := range add {
		t.imports.decl.Specs = append(t.imports.decl.Specs, &ast.ImportSpec{
			Path: stringLit(path),
		})
	}
	if len(t.imports.decl.Specs) == 0 {
		t.f.Decls = t.f.Decls[1:]
	}
}

func replace(n ast.Node, from, to ast.Expr) int {
	var cnt int
	switch n := n.(type) {
	case *ast.LabeledStmt:
		cnt += replace(n.Stmt, from, to)
	case *ast.IfStmt:
		if n.Cond == from {
			n.Cond = to
			cnt++
		}
		cnt += replace(n.Cond, from, to)
	case *ast.AssignStmt:
		if n.Rhs[0] == from {
			n.Rhs[0] = to
			cnt++
		}
		cnt += replace(n.Rhs[0], from, to)
	case *ast.ExprStmt:
		cnt += replace(n.X, from, to)
	case *ast.CallExpr:
		for i, a := range n.Args {
			if a == from {
				n.Args[i] = to
				cnt++
			}
			cnt += replace(a, from, to)
		}
		cnt += replace(n.Fun, from, to)
	case *ast.UnaryExpr:
		if n.X == from {
			n.X = to
			cnt++
		}
		cnt += replace(n.X, from, to)
	case *ast.BinaryExpr:
		if n.X == from {
			n.X = to
			cnt++
		}
		if n.Y == from {
			n.Y = to
			cnt++
		}
		cnt += replace(n.X, from, to)
		cnt += replace(n.Y, from, to)
	case *ast.SelectorExpr:
		if n.X == from {
			n.X = to
			cnt++
		}
		if n.Sel == from {
			n.Sel = to.(*ast.Ident)
			cnt++
		}
		cnt += replace(n.X, from, to)
	case *ast.CompositeLit:
		if n.Type == from {
			n.Type = to
			cnt++
		}
		for i, x := range n.Elts {
			if x == from {
				n.Elts[i] = to
				cnt++
			}
			cnt += replace(n.Elts[i], from, to)
		}
	case *ast.KeyValueExpr:
		if n.Key == from {
			n.Key = to
			cnt++
		}
		if n.Value == from {
			n.Value = to
			cnt++
		}
		cnt += replace(n.Key, from, to)
		cnt += replace(n.Value, from, to)
	}
	return cnt
}

func cntUsages(root ast.Node, id *ast.Ident) int {
	usages := 0
	ast.Inspect(root, func(n ast.Node) bool {
		if id == n {
			usages++
		}
		return true
	})
	return usages
}

func hasSideEffects(root ast.Node) bool {
	has := false
	ast.Inspect(root, func(n ast.Node) bool {
		if has {
			return false
		}
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return false
		}
		switch fnc := call.Fun.(type) {
		case *ast.SelectorExpr:
			switch fnc.Sel.Name {
			case "Atoi", "Itoa", "FormatFloat":
				// ok
			default:
				if fnc.Sel.Obj == nil {
					has = true
				} else if def, ok := fnc.Sel.Obj.Data.(*builtinDef); !ok || !def.Idemp {
					has = true
				}
			}
		default:
			has = true
		}
		return !has
	})
	return has
}

func arrayLen(x ast.Expr) (int, bool) {
	switch t := getType(x).(type) {
	case reflect.Type:
		if t.Kind() == reflect.Array {
			return t.Len(), true
		}
	case *TypeSet:
		if t.Def.Size > 1 {
			return t.Def.Size, true
		}
	}
	return 0, false
}

func getObj(x ast.Expr) *ast.Object {
	switch x := x.(type) {
	case *ast.Ident:
		return x.Obj
	case *ast.SelectorExpr:
		return x.Sel.Obj
	}
	return nil
}

func getType(x ast.Expr) any {
	switch x := x.(type) {
	case *ast.BasicLit:
		switch x.Kind {
		case token.STRING:
			return reflect.TypeOf("")
		}
	case *ast.IndexExpr:
		if obj := getObj(x.X); obj != nil {
			return obj.Type
		}
	}
	if obj := getObj(x); obj != nil {
		return obj.Type
	}
	return nil
}

func condition(x ast.Expr) ast.Expr {
	typeHint(x, reflBool)
	switch x := x.(type) {
	case *ast.BinaryExpr:
		switch x.Op {
		case token.EQL:
			if id, ok := x.Y.(*ast.Ident); ok {
				switch id.Name {
				case "true":
					return x.X
				case "false":
					return not(x.X)
				}
			}
		case token.NEQ:
			if id, ok := x.Y.(*ast.Ident); ok {
				switch id.Name {
				case "false":
					return x.X
				case "true":
					return not(x.X)
				}
			}
		}
	}
	return x
}
