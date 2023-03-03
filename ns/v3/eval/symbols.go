package eval

import "reflect"

//go:generate yaegi extract github.com/noxworld-dev/noxscript/ns/v3
//go:generate goimports -w .

var Symbols = make(map[string]map[string]reflect.Value)
