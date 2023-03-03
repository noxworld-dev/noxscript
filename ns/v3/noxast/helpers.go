package noxast

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"strconv"
)

func sel(obj ast.Expr, name string) ast.Expr {
	return &ast.SelectorExpr{X: obj, Sel: ast.NewIdent(name)}
}

func call(fnc ast.Expr, args ...ast.Expr) ast.Expr {
	return &ast.CallExpr{Fun: fnc, Args: args}
}

func callSel(obj ast.Expr, name string, args ...ast.Expr) ast.Expr {
	return &ast.CallExpr{Fun: sel(obj, name), Args: args}
}

func intLit(v int) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.INT, Value: strconv.FormatInt(int64(v), 10)}
}

func floatLit(v float32) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.FLOAT, Value: strconv.FormatFloat(float64(v), 'g', -1, 32)}
}

func stringLit(v string) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(v)}
}

func asInt(x ast.Expr) (int, bool) {
	switch x := x.(type) {
	case *ast.BasicLit:
		if x.Kind == token.INT {
			v, err := strconv.Atoi(x.Value)
			if err != nil {
				return 0, false
			}
			return v, true
		}
	}
	return 0, false
}

func asStr(x ast.Expr) (string, bool) {
	switch x := x.(type) {
	case *ast.BasicLit:
		if x.Kind == token.STRING {
			v, err := strconv.Unquote(x.Value)
			if err != nil {
				return "", false
			}
			return v, true
		}
	}
	return "", false
}

func takeAddr(x ast.Expr) ast.Expr {
	return &ast.UnaryExpr{Op: token.AND, X: x}
}

func not(x ast.Expr) ast.Expr {
	switch x := x.(type) {
	case *ast.UnaryExpr:
		if x.Op == token.NOT {
			return x.X
		}
	case *ast.BinaryExpr:
		op := x.Op
		switch op {
		case token.EQL:
			op = token.NEQ
		case token.NEQ:
			op = token.EQL
		default:
			op = 0
		}
		if op != 0 {
			return &ast.BinaryExpr{X: x.X, Op: op, Y: x.Y}
		}
	}
	return &ast.UnaryExpr{Op: token.NOT, X: x}
}

func printExpr(x ast.Expr) string {
	var buf bytes.Buffer
	format.Node(&buf, token.NewFileSet(), x)
	return buf.String()
}
